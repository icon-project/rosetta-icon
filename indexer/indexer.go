// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package indexer

import (
	"context"
	"errors"
	"fmt"
	"github.com/icon-project/goloop/common"
	"github.com/leeheonseung/rosetta-icon/configuration"
	"github.com/leeheonseung/rosetta-icon/icon/client_v1"
	"github.com/leeheonseung/rosetta-icon/services"
	"runtime"
	"time"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/storage/database"
	storageErrs "github.com/coinbase/rosetta-sdk-go/storage/errors"
	"github.com/coinbase/rosetta-sdk-go/storage/modules"
	"github.com/coinbase/rosetta-sdk-go/syncer"
	"github.com/coinbase/rosetta-sdk-go/types"
	sdkUtils "github.com/coinbase/rosetta-sdk-go/utils"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
)

const (
	// indexPlaceholder is provided to the syncer
	// to indicate we should both start from the
	// last synced block and that we should sync
	// blocks until exit (instead of stopping at
	// a particular height).
	indexPlaceholder = -1

	RetryDelay = 3 * time.Second

	nodeWaitSleep = 3 * time.Second

	// sizeMultiplier is used to multiply the memory
	// estimate for pre-fetching blocks. In other words,
	// this is the estimated memory overhead for each
	// block fetched by the indexer.
	sizeMultiplier = 5

	// zeroValue is 0 as a string
	zeroValue = "0"

	// overclockMultiplier is the amount
	// we multiply runtime.NumCPU by to determine
	// how many goroutines we should
	// spwan to handle block data sequencing.
	overclockMultiplier = 16

	// semaphoreWeight is the weight of each semaphore request.
	semaphoreWeight = int64(1)
)

// Client is used by the indexer to sync blocks.
type Client interface {
	NetworkStatus(context.Context) (*types.NetworkStatusResponse, error)
	GetBlock(context.Context, *client_v1.BlockRPCRequest) (*types.Block, error)
}

// Indexer caches blocks and provides balance query functionality.
type Indexer struct {
	cancel context.CancelFunc

	network *types.NetworkIdentifier

	client Client

	database       database.Database
	blockStorage   *modules.BlockStorage
	balanceStorage *modules.BalanceStorage
	workers        []modules.BlockWorker
}

// CloseDatabase closes a storage.Database. This should be called
// before exiting.
func (i *Indexer) CloseDatabase(ctx context.Context) {
	//todo logger
	err := i.database.Close(ctx)
	if err != nil {
	}

}

// defaultBadgerOptions returns a set of badger.Options optimized
// for running a Rosetta implementation.
func defaultBadgerOptions(
	dir string,
) badger.Options {
	opts := badger.DefaultOptions(dir)

	// By default, we do not compress the table at all. Doing so can
	// significantly increase memory usage.
	opts.Compression = options.None

	// Load tables into memory and memory map value logs.
	opts.TableLoadingMode = options.MemoryMap
	opts.ValueLogLoadingMode = options.MemoryMap

	// Use an extended table size for larger commits.
	opts.MaxTableSize = database.DefaultMaxTableSize

	// Smaller value log sizes means smaller contiguous memory allocations
	// and less RAM usage on cleanup.
	opts.ValueLogFileSize = database.DefaultLogValueSize

	// To allow writes at a faster speed, we create a new memtable as soon as
	// an existing memtable is filled up. This option determines how many
	// memtables should be kept in memory.
	opts.NumMemtables = 1

	// Don't keep multiple memtables in memory. With larger
	// memtable size, this explodes memory usage.
	opts.NumLevelZeroTables = 1
	opts.NumLevelZeroTablesStall = 2

	// This option will have a significant effect the memory. If the level is kept
	// in-memory, read are faster but the tables will be kept in memory. By default,
	// this is set to false.
	opts.KeepL0InMemory = false

	// We don't compact L0 on close as this can greatly delay shutdown time.
	opts.CompactL0OnClose = false

	// LoadBloomsOnOpen=false will improve the db startup speed. This is also
	// a waste to enable with a limited index cache size (as many of the loaded bloom
	// filters will be immediately discarded from the cache).
	opts.LoadBloomsOnOpen = false

	return opts
}

// Initialize returns a new Indexer.
func Initialize(
	ctx context.Context,
	config *configuration.Configuration,
	cancel context.CancelFunc,
	client Client,
) (*Indexer, error) {
	localStore, err := database.NewBadgerDatabase(
		ctx,
		config.IndexerPath,
		database.WithCustomSettings(defaultBadgerOptions(config.IndexerPath)),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: unable to initialize storage", err)
	}

	blockStorage := modules.NewBlockStorage(localStore, runtime.NumCPU()*overclockMultiplier)
	asserter, err := asserter.NewClientWithOptions(
		config.Network,
		config.GenesisBlockIdentifier,
		client_v1.OperationTypes,
		client_v1.OperationStatuses,
		services.Errors,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: unable to initialize asserter", err)
	}

	i := &Indexer{
		network:      config.Network,
		client:       client,
		cancel:       cancel,
		database:     localStore,
		blockStorage: blockStorage,
	}

	balanceStorage := modules.NewBalanceStorage(localStore)
	balanceStorage.Initialize(
		&BalanceStorageHelper{asserter},
		&BalanceStorageHandler{},
	)
	i.balanceStorage = balanceStorage

	i.workers = []modules.BlockWorker{balanceStorage}

	return i, nil
}

// waitForNode returns once bitcoind is ready to serve
// block queries.
func (i *Indexer) waitForNode(ctx context.Context) error {
	for {
		_, err := i.client.NetworkStatus(ctx)
		if err == nil {
			return nil
		}

		if err := sdkUtils.ContextSleep(ctx, RetryDelay); err != nil {
			return err
		}
	}
}

// Sync attempts to index Bitcoin blocks using
// the bitcoin.Client until stopped.
func (i *Indexer) Sync(ctx context.Context) error {
	if err := i.waitForNode(ctx); err != nil {
		return fmt.Errorf("%w: failed to wait for node", err)
	}

	i.blockStorage.Initialize(i.workers)

	startIndex := int64(indexPlaceholder)
	head, err := i.blockStorage.GetHeadBlockIdentifier(ctx)
	if err == nil {
		startIndex = head.Index + 1
	}

	// Load in previous blocks into syncer cache to handle reorgs.
	// If previously processed blocks exist in storage, they are fetched.
	// Otherwise, none are provided to the cache (the syncer will not attempt
	// a reorg if the cache is empty).
	pastBlocks := i.blockStorage.CreateBlockCache(ctx, syncer.DefaultPastBlockLimit)

	syncer := syncer.New(
		i.network,
		i,
		i,
		i.cancel,
		syncer.WithCacheSize(syncer.DefaultCacheSize),
		syncer.WithSizeMultiplier(sizeMultiplier),
		syncer.WithPastBlocks(pastBlocks),
	)

	return syncer.Sync(ctx, startIndex, indexPlaceholder)
}

// BlockAdded is called by the syncer when a block is added.
func (i *Indexer) BlockAdded(ctx context.Context, block *types.Block) error {

	err := i.blockStorage.AddBlock(ctx, block)
	if err != nil {
		return fmt.Errorf(
			"%w: unable to add block to storage %s:%d",
			err,
			block.BlockIdentifier.Hash,
			block.BlockIdentifier.Index,
		)
	}

	//logger.Debugw(
	//	"block added",
	//	"hash", block.BlockIdentifier.Hash,
	//	"index", block.BlockIdentifier.Index,
	//	"transactions", len(block.Transactions),
	//)

	return nil
}

// BlockSeen is called by the syncer when a block is encountered.
func (i *Indexer) BlockSeen(ctx context.Context, block *types.Block) error {

	// load intermediate
	err := i.blockStorage.SeeBlock(ctx, block)
	if err != nil {
		return fmt.Errorf(
			"%w: unable to encounter block to storage %s:%d",
			err,
			block.BlockIdentifier.Hash,
			block.BlockIdentifier.Index,
		)
	}

	return nil
}

// BlockRemoved is called by the syncer when a block is removed.
func (i *Indexer) BlockRemoved(
	ctx context.Context,
	blockIdentifier *types.BlockIdentifier,
) error {
	err := i.blockStorage.RemoveBlock(ctx, blockIdentifier)
	if err != nil {
		return fmt.Errorf(
			"%w: unable to remove block from storage %s:%d",
			err,
			blockIdentifier.Hash,
			blockIdentifier.Index,
		)
	}

	return nil
}

// NetworkStatus is called by the syncer to get the current
// network status.
func (i *Indexer) NetworkStatus(
	ctx context.Context,
	network *types.NetworkIdentifier,
) (*types.NetworkStatusResponse, error) {
	return i.client.NetworkStatus(ctx)
}

// Block is called by the syncer to fetch a block.
func (i *Indexer) Block(
	ctx context.Context,
	network *types.NetworkIdentifier,
	blockIdentifier *types.PartialBlockIdentifier,
) (*types.Block, error) {

	reqParams := &client_v1.BlockRPCRequest{}
	if blockIdentifier.Index != nil {
		reqParams = &client_v1.BlockRPCRequest{
			Height: common.HexInt64{Value: *blockIdentifier.Index}.String(),
		}
	} else if blockIdentifier.Hash != nil {
		reqParams = &client_v1.BlockRPCRequest{
			Hash: *blockIdentifier.Hash,
		}
	}
	block, err := i.client.GetBlock(ctx, reqParams)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetBlockTransaction returns a *types.Transaction if it is in the provided
// *types.BlockIdentifier.
func (i *Indexer) GetBlockTransaction(
	ctx context.Context,
	blockIdentifier *types.BlockIdentifier,
	transactionIdentifier *types.TransactionIdentifier,
) (*types.Transaction, error) {
	return i.blockStorage.GetBlockTransaction(
		ctx,
		blockIdentifier,
		transactionIdentifier,
	)
}

// GetBlockLazy returns a *types.BlockResponse from the indexer's block storage.
// All transactions in a block must be fetched individually.
func (i *Indexer) GetBlockLazy(
	ctx context.Context,
	blockIdentifier *types.PartialBlockIdentifier,
) (*types.BlockResponse, error) {
	return i.blockStorage.GetBlockLazy(ctx, blockIdentifier)
}

// GetBalance returns the balance of an account
// at a particular *types.PartialBlockIdentifier.
func (i *Indexer) GetBalance(
	ctx context.Context,
	accountIdentifier *types.AccountIdentifier,
	currency *types.Currency,
	blockIdentifier *types.PartialBlockIdentifier,
) (*types.Amount, *types.BlockIdentifier, error) {
	dbTx := i.database.ReadTransaction(ctx)
	defer dbTx.Discard(ctx)

	blockResponse, err := i.blockStorage.GetBlockLazyTransactional(
		ctx,
		blockIdentifier,
		dbTx,
	)
	if err != nil {
		return nil, nil, err
	}

	amount, err := i.balanceStorage.GetBalanceTransactional(
		ctx,
		dbTx,
		accountIdentifier,
		currency,
		blockResponse.Block.BlockIdentifier.Index,
	)
	if errors.Is(err, storageErrs.ErrAccountMissing) {
		return &types.Amount{
			Value:    zeroValue,
			Currency: currency,
		}, blockResponse.Block.BlockIdentifier, nil
	}
	if err != nil {
		return nil, nil, err
	}

	return amount, blockResponse.Block.BlockIdentifier, nil
}

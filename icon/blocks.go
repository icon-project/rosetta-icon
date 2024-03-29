package icon

import (
	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseGenesisBlock(blk *Block) (*types.Block, error) {
	transactions, _ := ParseGenesisTransaction(blk.Transactions)
	return &types.Block{
		BlockIdentifier: &types.BlockIdentifier{
			Index: blk.Number(),
			Hash:  blk.Hash(),
		},
		ParentBlockIdentifier: &types.BlockIdentifier{
			Index: blk.Number(),
			Hash:  blk.Hash(),
		},
		Timestamp:    blk.TimestampInMillis(),
		Transactions: transactions,
		Metadata:     blk.GenesisMeta(),
	}, nil
}

func ParseBlock(blk *Block) (*types.Block, error) {
	if blk.Number() == GenesisBlockIndex {
		return ParseGenesisBlock(blk)
	} else {
		transactions, _ := ParseTransactions(blk.Transactions)
		timestamp := blk.TimestampInMillis()
		if timestamp == 0 {
			timestamp = asserter.MinUnixEpoch
		}
		return &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: blk.Number(),
				Hash:  blk.Hash(),
			},
			ParentBlockIdentifier: &types.BlockIdentifier{
				Index: blk.Number() - 1,
				Hash:  blk.PrevHash(),
			},
			Timestamp:    timestamp,
			Transactions: transactions,
			Metadata:     blk.Meta(),
		}, nil
	}
}

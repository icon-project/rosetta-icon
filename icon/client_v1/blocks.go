package client_v1

import (
	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseGenesisBlock(blk *Block01a) (*types.Block, error) {
	transactions, _ := ParseGenesisTransaction(blk.Transactions)
	return &types.Block{
		BlockIdentifier: &types.BlockIdentifier{
			Index: blk.Number(),
			Hash:  blk.Hash(),
		},
		Timestamp:    blk.Time(),
		Transactions: transactions,
		Metadata:     blk.GenesisMeta(),
	}, nil
}

func ParseBlock01a(blk *Block01a) (*types.Block, error) {
	if blk.Number() == GenesisBlockIndex {
		return ParseGenesisBlock(blk)
	} else {
		transactions, _ := ParseTransactions(blk.Transactions)
		return &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: blk.Number(),
				Hash:  blk.Hash(),
			},
			ParentBlockIdentifier: &types.BlockIdentifier{
				Index: blk.Number() - 1,
				Hash:  blk.PrevHash(),
			},
			Timestamp:    blk.Time(),
			Transactions: transactions,
			Metadata:     blk.Meta(),
		}, nil
	}
}

func ParseBlock03(blk *Block03) (*types.Block, error) {
	transactions, _ := ParseTransactions(blk.Transactions)
	return &types.Block{
		BlockIdentifier: &types.BlockIdentifier{
			Index: blk.Number(),
			Hash:  blk.Hash(),
		},
		ParentBlockIdentifier: &types.BlockIdentifier{
			Index: blk.Number() - 1,
			Hash:  blk.PrevHash(),
		},
		Timestamp:    blk.Time(),
		Transactions: transactions,
		Metadata:     blk.Meta(),
	}, nil
}

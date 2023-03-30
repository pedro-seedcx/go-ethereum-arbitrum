package eth

import (
	"github.com/pedro-seedcx/go-ethereum-arbitrum/core"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/core/state"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/core/types"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/core/vm"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}

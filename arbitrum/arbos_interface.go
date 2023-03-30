package arbitrum

import (
	"context"

	"github.com/pedro-seedcx/go-ethereum-arbitrum/core"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}

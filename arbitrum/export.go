package arbitrum

import (
	"context"

	"github.com/pedro-seedcx/go-ethereum-arbitrum/common/hexutil"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/core"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/internal/ethapi"
	"github.com/pedro-seedcx/go-ethereum-arbitrum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}

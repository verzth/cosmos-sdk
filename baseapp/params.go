package baseapp

import (
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	sdk "github.com/verzth/cosmos-sdk/types"
)

// ParamStore defines the interface the parameter store used by the BaseApp must
// fulfill.
type ParamStore interface {
	Get(ctx sdk.Context) (*cmtproto.ConsensusParams, error)
	Has(ctx sdk.Context) bool
	Set(ctx sdk.Context, cp *cmtproto.ConsensusParams)
}

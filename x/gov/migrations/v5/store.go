package v5

import (
	storetypes "github.com/verzth/cosmos-sdk/store/types"

	"github.com/verzth/cosmos-sdk/codec"
	sdk "github.com/verzth/cosmos-sdk/types"
	v4 "github.com/verzth/cosmos-sdk/x/gov/migrations/v4"
	govv1 "github.com/verzth/cosmos-sdk/x/gov/types/v1"
)

// MigrateStore performs in-place store migrations from v4 (v0.47) to v5 (v0.48). The
// migration includes:
//
// Addition of the new proposal expedited parameters that are set to 0 by default.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	paramsBz := store.Get(v4.ParamsKey)

	var params govv1.Params
	cdc.MustUnmarshal(paramsBz, &params)

	defaultParams := govv1.DefaultParams()
	params.ExpeditedMinDeposit = defaultParams.ExpeditedMinDeposit
	params.ExpeditedVotingPeriod = defaultParams.ExpeditedVotingPeriod
	params.ExpeditedThreshold = defaultParams.ExpeditedThreshold
	params.ProposalCancelRatio = defaultParams.ProposalCancelRatio
	params.ProposalCancelDest = defaultParams.ProposalCancelDest

	bz, err := cdc.Marshal(&params)
	if err != nil {
		return err
	}

	store.Set(v4.ParamsKey, bz)

	return nil
}

package keeper

import (
	errorsmod "github.com/verzth/cosmos-sdk/errors"

	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/types/errors"
	"github.com/verzth/cosmos-sdk/x/crisis/types"
)

// GetConstantFee get's the constant fee from the store
func (k *Keeper) GetConstantFee(ctx sdk.Context) (constantFee sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ConstantFeeKey)
	if bz == nil {
		return constantFee
	}
	k.cdc.MustUnmarshal(bz, &constantFee)
	return constantFee
}

// GetConstantFee set's the constant fee in the store
func (k *Keeper) SetConstantFee(ctx sdk.Context, constantFee sdk.Coin) error {
	if !constantFee.IsValid() || constantFee.IsNegative() {
		return errorsmod.Wrapf(errors.ErrInvalidCoins, "negative or invalid constant fee: %s", constantFee)
	}

	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&constantFee)
	if err != nil {
		return err
	}

	store.Set(types.ConstantFeeKey, bz)
	return nil
}

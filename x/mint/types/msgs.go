package types

import (
	"github.com/verzth/cosmos-sdk/errors"

	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ sdk.Msg            = &MsgUpdateParams{}
	_ legacytx.LegacyMsg = &MsgUpdateParams{}
)

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}

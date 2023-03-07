package gov_test

import (
	"testing"

	"github.com/verzth/cosmos-sdk/math"
	"gotest.tools/v3/assert"

	sdk "github.com/verzth/cosmos-sdk/types"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	"github.com/verzth/cosmos-sdk/x/gov/types"
	v1 "github.com/verzth/cosmos-sdk/x/gov/types/v1"
	"github.com/verzth/cosmos-sdk/x/gov/types/v1beta1"
	stakingtypes "github.com/verzth/cosmos-sdk/x/staking/types"
)

var (
	valTokens           = sdk.TokensFromConsensusPower(42, sdk.DefaultPowerReduction)
	TestProposal        = v1beta1.NewTextProposal("Test", "description")
	TestDescription     = stakingtypes.NewDescription("T", "E", "S", "T", "Z")
	TestCommissionRates = stakingtypes.NewCommissionRates(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec())
)

// mkTestLegacyContent creates a MsgExecLegacyContent for testing purposes.
func mkTestLegacyContent(t *testing.T) *v1.MsgExecLegacyContent {
	msgContent, err := v1.NewLegacyContent(TestProposal, authtypes.NewModuleAddress(types.ModuleName).String())
	assert.NilError(t, err)

	return msgContent
}

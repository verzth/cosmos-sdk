package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/verzth/cosmos-sdk/codec"
	cryptoAmino "github.com/verzth/cosmos-sdk/crypto/codec"
	"github.com/verzth/cosmos-sdk/testutil/testdata"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/x/auth/migrations/legacytx"
	txtestutil "github.com/verzth/cosmos-sdk/x/auth/tx/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, txtestutil.NewTxConfigTestSuite(txGen))
}

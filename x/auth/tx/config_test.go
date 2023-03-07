package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/verzth/cosmos-sdk/codec"
	codectypes "github.com/verzth/cosmos-sdk/codec/types"
	"github.com/verzth/cosmos-sdk/std"
	"github.com/verzth/cosmos-sdk/testutil/testdata"
	sdk "github.com/verzth/cosmos-sdk/types"
	txtestutil "github.com/verzth/cosmos-sdk/x/auth/tx/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, txtestutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}

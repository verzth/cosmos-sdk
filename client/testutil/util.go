package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	runtimev1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/v1alpha1"
	"github.com/verzth/cosmos-sdk/codec"
	"github.com/verzth/cosmos-sdk/core/appconfig"
	"github.com/verzth/cosmos-sdk/depinject"
	_ "github.com/verzth/cosmos-sdk/runtime" // Register runtime module
)

var TestConfig = appconfig.Compose(&appv1alpha1.Config{
	Modules: []*appv1alpha1.ModuleConfig{
		{
			Name: "runtime",
			Config: appconfig.WrapAny(&runtimev1alpha1.Module{
				AppName: "clientTest",
			}),
		},
	},
})

func MakeTestCodec(t *testing.T) codec.Codec {
	var cdc codec.Codec
	err := depinject.Inject(TestConfig, &cdc)
	require.NoError(t, err)
	return cdc
}

package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/verzth/cosmos-sdk/client"
	"github.com/verzth/cosmos-sdk/client/flags"
	"github.com/verzth/cosmos-sdk/log"
	"github.com/verzth/cosmos-sdk/server"
	"github.com/verzth/cosmos-sdk/simapp"
	moduletestutil "github.com/verzth/cosmos-sdk/types/module/testutil"
	"github.com/verzth/cosmos-sdk/x/auth"
	banktypes "github.com/verzth/cosmos-sdk/x/bank/types"
	genutiltest "github.com/verzth/cosmos-sdk/x/genutil/client/testutil"
	genutiltypes "github.com/verzth/cosmos-sdk/x/genutil/types"
	"github.com/verzth/cosmos-sdk/x/staking"
)

func Test_TestnetCmd(t *testing.T) {
	home := t.TempDir()
	encodingConfig := moduletestutil.MakeTestEncodingConfig(staking.AppModuleBasic{}, auth.AppModuleBasic{})
	logger := log.NewNopLogger()
	cfg, err := genutiltest.CreateDefaultCometConfig(home)
	require.NoError(t, err)

	err = genutiltest.ExecInitCmd(simapp.ModuleBasics, home, encodingConfig.Codec)
	require.NoError(t, err)

	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithHomeDir(home).
		WithTxConfig(encodingConfig.TxConfig)

	ctx := context.Background()
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	cmd := testnetInitFilesCmd(simapp.ModuleBasics, banktypes.GenesisBalancesIterator{})
	cmd.SetArgs([]string{fmt.Sprintf("--%s=test", flags.FlagKeyringBackend), fmt.Sprintf("--output-dir=%s", home)})
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err)

	genFile := cfg.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	require.NoError(t, err)

	bankGenState := banktypes.GetGenesisStateFromAppState(encodingConfig.Codec, appState)
	require.NotEmpty(t, bankGenState.Supply.String())
}

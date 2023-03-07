package testutil

import (
	"context"
	"fmt"

	cmtcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/cli"
	"github.com/spf13/viper"

	"github.com/verzth/cosmos-sdk/log"

	"github.com/verzth/cosmos-sdk/client"
	"github.com/verzth/cosmos-sdk/codec"
	"github.com/verzth/cosmos-sdk/server"
	"github.com/verzth/cosmos-sdk/testutil"
	"github.com/verzth/cosmos-sdk/types/module"
	genutilcli "github.com/verzth/cosmos-sdk/x/genutil/client/cli"
)

func ExecInitCmd(testMbm module.BasicManager, home string, cdc codec.Codec) error {
	logger := log.NewNopLogger()
	cfg, err := CreateDefaultCometConfig(home)
	if err != nil {
		return err
	}

	cmd := genutilcli.InitCmd(testMbm, home)
	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.WithCodec(cdc).WithHomeDir(home)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

	cmd.SetArgs([]string{"appnode-test", fmt.Sprintf("--%s=%s", cli.HomeFlag, home)})

	return cmd.ExecuteContext(ctx)
}

func CreateDefaultCometConfig(rootDir string) (*cmtcfg.Config, error) {
	conf := cmtcfg.DefaultConfig()
	conf.SetRoot(rootDir)
	cmtcfg.EnsureRoot(rootDir)

	if err := conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %v", err)
	}

	return conf, nil
}

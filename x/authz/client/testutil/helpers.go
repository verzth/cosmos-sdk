package authz

import (
	"github.com/verzth/cosmos-sdk/client"
	"github.com/verzth/cosmos-sdk/testutil"
	clitestutil "github.com/verzth/cosmos-sdk/testutil/cli"
	"github.com/verzth/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}

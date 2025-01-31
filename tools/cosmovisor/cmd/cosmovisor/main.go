package main

import (
	"context"
	"os"

	"github.com/verzth/cosmos-sdk/log"
	cverrors "github.com/verzth/cosmos-sdk/tools/cosmovisor/errors"
)

func main() {
	logger := log.NewLoggerWithKV(os.Stdout, log.ModuleKey, "cosmovisor")
	ctx := context.WithValue(context.Background(), log.ContextKey, logger)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		cverrors.LogErrors(logger, "", err)
		os.Exit(1)
	}
}

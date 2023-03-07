package main

import (
	"os"

	"github.com/verzth/cosmos-sdk/server"
	svrcmd "github.com/verzth/cosmos-sdk/server/cmd"
	"github.com/verzth/cosmos-sdk/simapp"
	"github.com/verzth/cosmos-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

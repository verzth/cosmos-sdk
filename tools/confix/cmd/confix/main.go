package main

import (
	"os"

	confixcmd "github.com/verzth/cosmos-sdk/tools/confix/cmd"
)

func main() {
	if err := confixcmd.ConfigCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

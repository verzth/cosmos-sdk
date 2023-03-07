package main

import (
	"github.com/verzth/cosmos-sdk/tools/hubl/internal"
)

func main() {
	cmd, err := internal.RootCommand()
	if err != nil {
		panic(err)
	}

	if err = cmd.Execute(); err != nil {
		panic(err)
	}
}

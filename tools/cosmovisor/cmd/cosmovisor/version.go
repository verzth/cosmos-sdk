package main

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/verzth/cosmos-sdk/log"
)

func init() {
	versionCmd.Flags().StringP(OutputFlag, "o", "text", "Output format (text|json)")
	rootCmd.AddCommand(versionCmd)
}

// OutputFlag defines the output format flag
var OutputFlag = "output"

var versionCmd = &cobra.Command{
	Use:          "version",
	Short:        "Prints the version of Cosmovisor.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := cmd.Context().Value(log.ContextKey).(log.Logger)

		if val, err := cmd.Flags().GetString(OutputFlag); val == "json" && err == nil {
			return printVersionJSON(logger, args)
		}

		return printVersion(logger, args)
	},
}

func getVersion() string {
	version, ok := debug.ReadBuildInfo()
	if !ok {
		panic("failed to get cosmovisor version")
	}

	return strings.TrimSpace(version.Main.Version)
}

func printVersion(logger log.Logger, args []string) error {
	fmt.Printf("cosmovisor version: %s\n", getVersion())

	if err := Run(logger, append([]string{"version"}, args...)); err != nil {
		return fmt.Errorf("failed to run version command: %w", err)
	}

	return nil
}

func printVersionJSON(logger log.Logger, args []string) error {
	buf := new(strings.Builder)

	// disable logger
	zl := logger.Impl().(*zerolog.Logger)
	logger = log.NewCustomLogger(zl.Level(zerolog.Disabled))

	if err := Run(
		logger,
		[]string{"version", "--long", "--output", "json"},
		StdOutRunOption(buf),
	); err != nil {
		return fmt.Errorf("failed to run version command: %w", err)
	}

	out, err := json.Marshal(struct {
		Version    string          `json:"cosmovisor_version"`
		AppVersion json.RawMessage `json:"app_version"`
	}{
		Version:    getVersion(),
		AppVersion: json.RawMessage(buf.String()),
	})
	if err != nil {
		return fmt.Errorf("can't print version output, expected valid json from APP, got: %s - %w", buf.String(), err)
	}

	fmt.Println(string(out))
	return nil
}

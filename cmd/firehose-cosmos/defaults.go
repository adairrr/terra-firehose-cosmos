package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamingfast/dlauncher/launcher"
)

var (
	// Data dir (for local operations only)
	DataDir = "./fh-data"

	// GRPC Service Addresses
	BlockStreamServingAddr  = "0.0.0.0:9000"
	RelayerServingAddr      = "0.0.0.0:9010"
	MergerServingAddr       = "0.0.0.0:9020"
	FirehoseGRPCServingAddr = "0.0.0.0:9030"

	// Blocks store
	MergedBlocksStoreURL string = "file://{fh-data-dir}/storage/merged-blocks"
	OneBlockStoreURL     string = "file://{fh-data-dir}/storage/one-blocks"

	// Protocol defaults
	FirstStreamableBlock uint64 = 0
)

func init() {
	launcher.RegisterCommonFlags = func(cmd *cobra.Command) error {
		initCommonFlags(cmd.Flags())
		return nil
	}
}

func initCommonFlags(flags *pflag.FlagSet) {
	// Logging
	flags.IntP("verbose", "v", 3, "Enables verbose output (-vvvv for max verbosity)")
	flags.String("log-format", "text", "Logging format")
	flags.StringP("config", "c", "firehose.yml", "Configuration file for the firehose")
	flags.StringP("data-dir", "d", DataDir, "Path to data storage for all components of firehose")

	// Common stores configuration flags
	flags.String("common-blocks-store-url", MergedBlocksStoreURL, "Store URL (with prefix) where to read/write")
	flags.String("common-oneblock-store-url", OneBlockStoreURL, "Store URL (with prefix) to read/write one-block files")
	flags.String("common-blockstream-addr", RelayerServingAddr, "GRPC endpoint to get real-time blocks")
	flags.Uint64("common-first-streamable-block", FirstStreamableBlock, "First streamable block number")

	// Authentication, metering and rate limiter plugins
	flags.String("common-auth-plugin", "null://", "Auth plugin URI, see streamingfast/dauth repository")
	flags.String("common-metering-plugin", "null://", "Metering plugin URI, see streamingfast/dmetering repository")

	// System Behavior
	flags.Duration("common-startup-delay", 0, "Delay before launching firehose process")
	flags.Duration("common-shutdown-delay", 5, "Add a delay between receiving SIGTERM signal and shutting down apps. Apps will respond negatively to /healthz during this period")
}

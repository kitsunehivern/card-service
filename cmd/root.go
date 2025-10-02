package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	grpcAddr string
	httpAddr string
)

var rootCmd = &cobra.Command{
	Use:   "model-service",
	Short: "Card Service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {}

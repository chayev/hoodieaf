package cmd

import (
	"fmt"
	"os"

	"github.com/chayev/hoodieaf/lib"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "hoodieaf <city>",
	Short:   "Find out if you need a hoodie or not.",
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(lib.Hoodie(args[0]))
		// lib.Hoodie(args[0])
	},
}

// Execute returns usage guide
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

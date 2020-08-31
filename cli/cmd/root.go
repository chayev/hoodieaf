package cmd

import (
	"fmt"
	"os"

	"github.com/chayev/hoodieaf/lib"

	"github.com/spf13/cobra"
)

var cfgFile string

var (
	search string

	rootCmd = &cobra.Command{
		Use:     "hoodieaf",
		Short:   "Find out if you need a hoodie or not.",
		Version: "v0.1.0",
		Run: func(cmd *cobra.Command, args []string) {

			if search != "" {
				var temp = lib.GetTemp(search)
				fmt.Printf("Temp in %s: %d \n", search, temp)
				if temp > 15 {
					fmt.Println("no hoodie")
				} else {
					fmt.Println("hoodie")
				}
			} else {
				fmt.Println("Enter `hoodieaf -h` for help.")
			}
		},
	}
)

// Execute returns usage guide
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVarP(&search, "search", "s", "", "name of city you would like to check")

}

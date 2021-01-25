package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch raw data into file",
	Long:  `fetch raw data into file if not existes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetch args:", codeArray)
	},
}

package cmd

import (
	"github.com/ping40/up55/analyse"
	"github.com/ping40/up55/fetch"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(analyseCmd)
}

var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "analyse files into human readable info",
	Long:  `analyse files into human readable info`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range codeArray {
			fetch.Fetch(v)
			analyse.Anaylse(v)
		}
	},
}

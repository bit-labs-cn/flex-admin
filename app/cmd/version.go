package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
	},
}

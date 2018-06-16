package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "jisoo2",
		Short: "hellow welcome to jisoo clean links for webs",

		SilenceUsage: true,
	}
	cmd.AddCommand(getLinks())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

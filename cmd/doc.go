package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func newDocCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "doc",
		Short:  "Generate documentation",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := os.MkdirAll("./docs", 0755)
			if err != nil {
				return err
			}

			return doc.GenMarkdownTree(cmd.Parent(), "./docs")
		},
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// rootCmd represents the base command when called without any subcommands.
	rootCmd := &cobra.Command{
		Use:     "gh-check",
		Short:   "Github check is a tool that help you to check issue/PR",
		Long:    "Github check is a tool that help you to check issue/PR",
		Version: version,
	}

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newDocCmd())
	rootCmd.AddCommand(newLabelsCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

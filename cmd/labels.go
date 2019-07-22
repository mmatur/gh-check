package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/mmatur/gh-check/labels"
	"github.com/spf13/cobra"
)

func newLabelsCmd() *cobra.Command {
	cfg := labels.Config{}

	labelsCmd := &cobra.Command{
		Use:   "labels",
		Short: "Check if issue contains labels.",
		Long:  "Check if issue contains labels.",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return validateRequiredFlags(cfg)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := labels.Labels(&cfg); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}

	labelsFlags := labelsCmd.Flags()
	labelsFlags.StringVar(&cfg.Owner, "owner", "", "Repository owner.")
	labelsFlags.StringVar(&cfg.Name, "name", "", "Repository name.")
	labelsFlags.StringVar(&cfg.GithubToken, "github-token", os.Getenv("GITHUB_TOKEN"), "The github token to used.")
	labelsFlags.IntVar(&cfg.Number, "number", 0, "Issue/PR number.")
	labelsFlags.StringSliceVar(&cfg.Labels, "labels", nil, "Labels need to be present on the issue.")
	labelsFlags.BoolVar(&cfg.Debug, "debug", false, "Enable debug log.")

	return labelsCmd
}

func validateRequiredFlags(cfg labels.Config) error {
	var missingFlagNames []string

	value := reflect.ValueOf(cfg)
	for i := 0; i < value.NumField(); i++ {
		switch value.Field(i).Kind() {
		case reflect.String:
			if value.Field(i).String() != "" {
				continue
			}
		case reflect.Int8:
			if value.Field(i).Int() != 0 {
				continue
			}
		case reflect.Slice:
			if value.Field(i).Len() > 0 {
				continue
			}
		default:
			if value.Field(i).String() != "" {
				continue
			}
		}

		fieldType := value.Type().Field(i)
		name := fieldType.Tag.Get("flag")
		if name == "" {
			name = fieldType.Name
		}

		missingFlagNames = append(missingFlagNames, name)
	}

	if len(missingFlagNames) > 0 {
		return fmt.Errorf(`required flag(s) "%s" not set`, strings.Join(missingFlagNames, `", "`))
	}

	return nil
}

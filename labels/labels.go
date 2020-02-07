package labels

import (
	"context"
	"log"

	"github.com/mmatur/gh-check/gh"
)

// Config is the labels configuration.
type Config struct {
	Owner       string   `flag:"owner"`
	Name        string   `flag:"name"`
	GithubToken string   `flag:"github-token"`
	Number      int      `flag:"number"`
	Labels      []string `flag:"labels"`
	Debug       bool     `flag:"debug"`
}

// Labels checks are all present on the issue.
func Labels(cfg *Config) error {
	ctx := context.Background()
	if cfg.Debug {
		log.Println("Creating Github client")
	}
	client := gh.NewGitHubClient(ctx, cfg.GithubToken)
	ghub := gh.NewGHub(client, cfg.Debug)

	return ghub.HasLabels(ctx, cfg.Owner, cfg.Name, cfg.Number, cfg.Labels)
}

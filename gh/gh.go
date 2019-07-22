package gh

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GHub GitHub helper.
type GHub struct {
	client *github.Client
	debug  bool
}

// NewGHub create a new GHub
func NewGHub(client *github.Client, debug bool) *GHub {
	return &GHub{client: client, debug: debug}
}

// HasLabels checks if issue/PR has labels.
func (g *GHub) HasLabels(ctx context.Context, owner string, name string, number int, labels []string) error {
	if g.debug {
		log.Printf("Retrieving issue %q in %s/%s\n", number, owner, name)
	}
	issue, _, err := g.client.Issues.Get(ctx, owner, name, number)
	if err != nil {
		return err
	}

	if g.debug {
		log.Printf("%s\n", issue.String())
	}

	var issueLabels []string
	for _, lbl := range issue.Labels {
		issueLabels = append(issueLabels, lbl.GetName())
	}

	for _, value := range labels {
		if !contains(issueLabels, value) {
			return fmt.Errorf("unable to find label %q, in %s", value, issueLabels)
		}
	}

	return nil
}

// NewGitHubClient create a new GitHub client
func NewGitHubClient(ctx context.Context, token string, gitHubURL *url.URL) *github.Client {
	var client *github.Client
	if len(token) == 0 {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	if gitHubURL != nil {
		client.BaseURL = gitHubURL
	}

	return client
}

func contains(values []string, value string) bool {
	for _, val := range values {
		if value == val {
			return true
		}
	}
	return false
}

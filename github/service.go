package github

import (
	"fmt"

	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v50/github"
	"github.com/spf13/viper"
)

type GithubEventService struct {
	Handler *githubevents.EventHandler
}

func NewGithubEventService() *GithubEventService {
	service := &GithubEventService{
		Handler: githubevents.New(viper.GetString("github.token")),
	}
	service.Handler.OnIssueCommentCreated(
		func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
			fmt.Printf("%s made a comment!", event.Sender.Login)
			return nil
		},
	)
	return service
}

package github

import "github.com/cbrgm/githubevents/githubevents"

type GithubEventService struct {
	handler *githubevents.EventHandler
}

func NewGithubEventService() *GithubEventService {
	return &GithubEventService{
		handler: githubevents.New("webhookSecret"),
	}
}



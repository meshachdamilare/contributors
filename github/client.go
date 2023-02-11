package github

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	token  string
	client http.Client
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("Please provide your token")
	}
	return &Client{
		token: token,
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}

func (c *Contributor) ContributionList(repo) ([]Contributor, error) {

}

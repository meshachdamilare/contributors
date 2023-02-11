package github

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Client) ContributionList(repo string) ([]Contributor, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contributors", repo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server returns %d error", res.StatusCode)
	}

	var cons []Contributor
	err = json.NewDecoder(res.Body).Decode(&cons)
	if err != nil {
		return nil, err
	}
	return cons, nil
}

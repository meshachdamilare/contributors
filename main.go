package main

import (
	"fmt"
	"github.com/meshachdamilare/contributors/github"
	"log"
	"os"
)

type ContributorLister interface {
	ContributionList(repo string) ([]github.Contributor, error)
}

func main() {
	token := os.Getenv("token")
	if token == "" {
		log.Print("Token not found. You must set it in your environment like")
		log.Print("export GITHUB_TOKEN=000a0aaaa0000a00000000aaa00000000a000000")
		log.Print("You can generate a token at https://github.com/settings/tokens")
		os.Exit(1)
	}

	repo := "AltGophers/numbase"

	client, err := github.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}
	if err := process(client, repo); err != nil {
		log.Fatal(err)
	}

}

func process(c ContributorLister, repo string) error {
	cons, err := c.ContributionList(repo)
	if err != nil {
		return err
	}
	for _, ct := range cons {
		fmt.Println("List of Repo contributor are: ")
		fmt.Printf("Name: %s\\ttGithub Url: %s\t\tContributions: %d\n", ct.Login, ct.Url, ct.Contributions)
	}
	return nil
}

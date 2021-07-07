package main

import (
	"fmt"
	"github.com/google/go-github/v36/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatalln(err)
	}
	for i := range(repos) {
		fmt.Println(*repos[i].Name)
	}
}

package main

import (
	"fmt"
	"strings"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/new_yearify/pkg/api"
	"github.com/gleich/new_yearify/pkg/out"
)

func main() {
	PAT := out.Ask("What is your PAT (personal access token)?")
	if PAT == "" || !strings.HasPrefix(PAT, "ghp_") {
		lumber.FatalMsg("Please enter a valid response")
	}

	client := api.Client(PAT)

	username, err := api.Username(client)
	if err != nil {
		lumber.Fatal(err, "Failed to get user's username")
	}

	repos, err := api.Repos(username, client)
	if err != nil {
		lumber.Fatal(err, "Failed to load repos")
	}
	fmt.Println(repos)
}

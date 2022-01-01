package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/new_yearify/pkg/api"
	"github.com/gleich/new_yearify/pkg/out"
	"github.com/gleich/new_yearify/pkg/update"
)

func main() {
	PAT := out.Ask("What is your PAT (personal access token)?")
	if PAT == "" || !strings.HasPrefix(PAT, "ghp_") {
		lumber.FatalMsg("Please enter a valid personal access token")
	}

	client := api.Client(PAT)

	repos, err := api.Repos(client)
	if err != nil {
		lumber.Fatal(err, "Failed to load repos")
	}

	tmpDir, err := update.CreateTmpDir()
	if err != nil {
		lumber.Fatal(err, "Failed to create temp directory for cloning")
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		lumber.Fatal(err, "Failed to change directory to temporary directory for cloning")
	}

	updates := 0
	for i, repo := range repos {
		if repo.IsArchived || repo.IsDisabled || repo.IsEmpty || repo.IsFork || repo.IsMirror {
			continue
		}
		err = update.Clone(repo)
		if err != nil {
			lumber.Fatal(err, "Failed to clone", repo.NameWithOwner)
		}
		lumber.Success("Cloned", repo.NameWithOwner, fmt.Sprintf("(%v/%v)", i+1, len(repos)))

		updated, err := update.Copyright(repo)
		if err != nil {
			lumber.Fatal(err, "Failed to update copyright for", repo.NameWithOwner)
		}

		if updated {
			updates++
			err = update.Commit(repo)
			if err != nil {
				lumber.Fatal(err, "Failed to commit & push changes for", repo.NameWithOwner)
			}
		}

		err = os.Chdir("..")
		if err != nil {
			lumber.Fatal(err, "Failed to change directory up out of the repository")
		}
	}

	fmt.Println()
	lumber.Success("Updated", updates, "repositories from", os.Args[1], "to", os.Args[2])
}

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
		lumber.FatalMsg("Please enter a valid response")
	}

	client := api.Client(PAT)

	repos, err := api.Repos(client)
	if err != nil {
		lumber.Fatal(err, "Failed to load repos")
	}
	fmt.Println(repos)

	tmpDir, err := update.CreateTmpDir()
	if err != nil {
		lumber.Fatal(err, "Failed to create temp directory for cloning")
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		lumber.Fatal(err, "Failed to change directory to temporary directory for cloning")
	}

	for _, repo := range repos {
		if repo.IsArchived || repo.IsDisabled || repo.IsEmpty || repo.IsFork || repo.IsMirror {
			continue
		}
		// loc := filepath.Join(tmpDir, repo.Name)
		err = update.Clone(repo)
		if err != nil {
			lumber.Fatal(err, "Failed to clone", repo.NameWithOwner)
		}
		break
	}
}

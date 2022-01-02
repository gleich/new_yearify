package update

import (
	"os"
	"os/exec"

	"github.com/gleich/new_yearify/pkg/api"
)

func Clone(repo api.Repo) error {
	source := repo.URL + ".git"

	// switch to ssh if flag "--ssh" is passed in
	for _, arg := range os.Args {
		if arg == "--ssh" {
			source = "git@github.com:" + repo.NameWithOwner
			break
		}
	}

	err := exec.Command("git", "clone", source).Run()
	if err != nil {
		return err
	}
	return nil
}

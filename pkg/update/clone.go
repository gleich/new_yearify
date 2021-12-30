package update

import (
	"os/exec"

	"github.com/gleich/new_yearify/pkg/api"
)

func Clone(repo api.Repo) error {
	err := exec.Command("git", "clone", repo.URL+".git").Run()
	if err != nil {
		return err
	}
	return nil
}

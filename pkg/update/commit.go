package update

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/new_yearify/pkg/api"
)

func Commit(repo api.Repo) error {
	err := exec.Command("git", "add", ".").Run()
	if err != nil {
		return err
	}

	err = exec.Command("git", "commit", "-m", fmt.Sprintf("%v -> %v", os.Args[1], os.Args[2])).Run()
	if err != nil {
		return err
	}

	lumber.Info("Updated", repo.NameWithOwner, "to new year")
	return nil
}

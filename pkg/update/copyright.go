package update

import (
	"os"
	"strings"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/new_yearify/pkg/api"
)

func Copyright(repo api.Repo) error {
	err := os.Chdir(repo.Name)
	if err != nil {
		return err
	}

	fsObjects, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, fsObject := range fsObjects {
		name := fsObject.Name()
		if !fsObject.IsDir() && strings.Contains(strings.ToLower(name), "license") {
			b, err := os.ReadFile(name)
			if err != nil {
				return err
			}

			content := string(b)
			patchedFile := strings.ReplaceAll(content, os.Args[1], os.Args[2])
			if content != patchedFile {
				lumber.Debug(name)
				err = os.WriteFile(name, []byte(patchedFile), fsObject.Type().Perm())
				if err != nil {
					return err
				}
			}
		}
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}

package update

import (
	"os"
	"path/filepath"
)

func CreateTmpDir() (string, error) {
	loc := filepath.Join(os.TempDir(), "new_yearify")
	if _, err := os.Stat(loc); err != nil {
		err := os.Mkdir(loc, 0755)
		if err != nil {
			return "", err
		}
	}
	return loc, nil
}

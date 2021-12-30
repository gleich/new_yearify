package api

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func Username(client *githubv4.Client) (string, error) {
	var query struct {
		Viewer struct {
			Login string
		}
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		return "", err
	}
	return query.Viewer.Login, nil
}

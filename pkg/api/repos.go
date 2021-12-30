package api

import (
	"context"
	"fmt"

	"github.com/gleich/lumber/v2"
	"github.com/shurcooL/githubv4"
)

type Repo struct {
	Name     string
	CloneURL string
}

func Repos(username string, client *githubv4.Client) ([]Repo, error) {
	fmt.Println()
	lumber.Info("Fetching repos")

	repos := []Repo{}
	vars := map[string]interface{}{
		"cursor": (*githubv4.String)(nil),
		"login":  githubv4.String(username),
	}
	for {
		query := struct {
			User struct {
				Repositories struct {
					Nodes []struct {
						URL  string
						Name string
					}
					PageInfo struct {
						EndCursor   string
						HasNextPage bool
					}
				} `graphql:"repositories(isFork: false, isLocked: false, first: 100, after: $cursor)"`
			} `graphql:"user(login: $login)"`
		}{}

		err := client.Query(context.Background(), &query, vars)
		if err != nil {
			return []Repo{}, err
		}

		for _, repo := range query.User.Repositories.Nodes {
			repos = append(repos, Repo{Name: repo.Name, CloneURL: repo.URL + ".git"})
		}
		if !query.User.Repositories.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.String(query.User.Repositories.PageInfo.EndCursor)
	}
	lumber.Success("Loaded", len(repos), "repos")
	return repos, nil
}

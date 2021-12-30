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
	fmt.Println(client)
}

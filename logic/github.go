package logic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bfv/pasoe-cli/model"
)

func GetLatestRelease() (model.GitRelease, error) {

	releases, err := GetReleases()
	if err != nil {
		fmt.Println("err3")
		printError(err)
		os.Exit(1)
	}

	return releases[0], nil
}

func GetReleases() (model.GitReleases, error) {

	res, _ := doGithubRequest("GET", "releases")
	defer res.Body.Close()

	releases := model.GitReleases{}
	err := json.NewDecoder(res.Body).Decode(&releases)
	if err != nil {
		fmt.Println("err3")
		printError(err)
		os.Exit(1)
	}

	return releases, nil
}

func doGithubRequest(verb string, path string) (*http.Response, error) {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf("https://api.github.com/repos/bfv/pasoe-cli/%v", path)
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		fmt.Println("err1")
		printError(err)
		os.Exit(1)
	}

	req.Header.Add("Accept", "application/vnd.github+json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err2")
		printError(err)
		os.Exit(1)
	}

	return res, nil
}

package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type _tagCommit struct {
	Sha string `json:"sha"`
}

type _tag struct {
	Name   string     `json:"name"`
	Commit _tagCommit `json:"commit"`
}

type Tag struct {
	Version string
	Sha     string
}

func fetchRepositoryTags(repo string) ([]_tag, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/tags", repo))
	if err != nil {
		return nil, fmt.Errorf("an error occured while fetching %s repo tags", repo)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	var githubTags []_tag
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return githubTags, json.Unmarshal(body, &githubTags)
}

func GetLatestRepositoryTag(repo string) (*Tag, error) {
	log.Println("> Fetching latest version..")

	tags, err := fetchRepositoryTags(repo)
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		if !strings.HasPrefix(tag.Name, "v1.") {
			continue
		}

		if split := strings.Split(tag.Name, "."); len(split) > 0 {
			return &Tag{
				Version: split[len(split)-1],
				Sha:     tag.Commit.Sha,
			}, nil
		}
	}

	return nil, fmt.Errorf("no tags found for the repo %s", repo)
}

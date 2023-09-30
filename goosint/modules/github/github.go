package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Item struct {
	Commit struct {
		Author struct {
			Email    string `json:"email"`
			Username string `json:"name"`
		} `json:"author"`
		Committer struct {
			Email    string `json:"email"`
			Username string `json:"name"`
		} `json:"committer"`
	} `json:"commit"`
}

type APIResponse struct {
	TotalCount int    `json:"total_count"`
	Items      []Item `json:"items"`
}

func GithubSearchEmail(email string) []string {
	apiURL1 := "https://api.github.com/search/commits?q=author-email:"
	apiURL2 := "https://api.github.com/search/commits?q=committer-email:"
	var allCommits []Item

	page := 1
	perPage := 100
	totalCount := 0

	for {
		url1 := fmt.Sprintf("%s%s&per_page=%d&page=%d", apiURL1, email, perPage, page)
		url2 := fmt.Sprintf("%s%s&per_page=%d&page=%d", apiURL2, email, perPage, page)

		// Query API 1
		response1, err := http.Get(url1)
		if err != nil {
			log.Fatal(err)
		}
		defer response1.Body.Close()

		body1, err := io.ReadAll(response1.Body)
		if err != nil {
			log.Fatal(err)
		}

		var apiResponse1 APIResponse
		err = json.Unmarshal(body1, &apiResponse1)
		if err != nil {
			log.Fatal(err)
		}

		// Query API 2
		response2, err := http.Get(url2)
		if err != nil {
			log.Fatal(err)
		}
		defer response2.Body.Close()

		body2, err := io.ReadAll(response2.Body)
		if err != nil {
			log.Fatal(err)
		}

		var apiResponse2 APIResponse
		err = json.Unmarshal(body2, &apiResponse2)
		if err != nil {
			log.Fatal(err)
		}

		totalCount = apiResponse1.TotalCount + apiResponse2.TotalCount

		allCommits = append(allCommits, apiResponse1.Items...)
		allCommits = append(allCommits, apiResponse2.Items...)

		if len(allCommits) >= totalCount {
			break
		}

		page++
		time.Sleep(1000 * time.Millisecond)
	}

	usernameResult := []string{}
	for _, commit := range allCommits {
		usernameResult = append(usernameResult, commit.Commit.Author.Username)
		usernameResult = append(usernameResult, commit.Commit.Committer.Username)
	}
	usernameResult = unique(usernameResult)
	return usernameResult
}

func GithubSearchName(username string) []string {
	apiURL1 := "https://api.github.com/search/commits?q=committer-name:"
	apiURL2 := "https://api.github.com/search/commits?q=author-name:"
	var allCommits []Item

	page := 1
	perPage := 100
	totalCount := 0

	for {
		url1 := fmt.Sprintf("%s%s&per_page=%d&page=%d", apiURL1, username, perPage, page)
		url2 := fmt.Sprintf("%s%s&per_page=%d&page=%d", apiURL2, username, perPage, page)

		// Query API 1
		response1, err := http.Get(url1)
		check(err)
		defer response1.Body.Close()

		body1, err := io.ReadAll(response1.Body)
		check(err)

		var apiResponse1 APIResponse
		err = json.Unmarshal(body1, &apiResponse1)
		check(err)

		// Query API 2
		response2, err := http.Get(url2)
		check(err)
		defer response2.Body.Close()

		body2, err := io.ReadAll(response2.Body)
		check(err)

		var apiResponse2 APIResponse
		err = json.Unmarshal(body2, &apiResponse2)
		check(err)

		totalCount = apiResponse1.TotalCount + apiResponse2.TotalCount

		allCommits = append(allCommits, apiResponse1.Items...)
		allCommits = append(allCommits, apiResponse2.Items...)

		if len(allCommits) >= totalCount {
			break
		}

		page++
		time.Sleep(1000 * time.Millisecond)
	}

	emailResult := []string{}
	for _, commit := range allCommits {
		emailResult = append(emailResult, commit.Commit.Author.Email)
		emailResult = append(emailResult, commit.Commit.Committer.Email)
	}
	emailResult = unique(emailResult)
	return emailResult
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

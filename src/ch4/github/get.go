package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetIssues(owner, repo string) (*[]Issue, error) {

	resp, err := http.Get(ReposURL + "/" + owner + "/" + repo + "/issues")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get issues failed: %s\n", resp.Status)
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}

func GetIssue(owner, repo string, number int) (*Issue, error) {

	resp, err := http.Get(ReposURL + "/" + owner + "/" + repo + "/issues/" + strconv.Itoa(number))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get issue %v failed: %s\n", number, resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}

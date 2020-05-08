package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// TODO: these have fairly similar sending procedures, can be grouped

func CreateIssue(owner, repo string, issue *NewIssue) (*Issue, error) {

	url := ReposURL + "/" + owner + "/" + repo + "/issues"
	body, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	token := os.Getenv(APITokenEnvVar)
	req.Header.Add("Authorization", "token "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Post issue failed: %s\n", res.Status)
	}

	var result Issue
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}

func CloseIssue(owner, repo string, number int) (*Issue, error) {
	url := ReposURL + "/" + owner + "/" + repo + "/issues/" + strconv.Itoa(number)
	closedState := &IssueState{"closed"}

	body, err := json.Marshal(closedState)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	token := os.Getenv(APITokenEnvVar)
	req.Header.Add("Authorization", "token "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Close issue failed: %s\n", res.Status)
	}

	var result Issue
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}

func UpdateIssue(owner, repo string, number int) (*Issue, error) {
	url := ReposURL + "/" + owner + "/" + repo + "/issues/" + strconv.Itoa(number)
	update := &IssueUpdate{}
	// TODO: add fields

	body, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	token := os.Getenv(APITokenEnvVar)
	req.Header.Add("Authorization", "token "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Update issue failed: %s\n", res.Status)
	}

	var result Issue
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}

package main

import (
	"ch4/github" // auth is done by placing token in env var GITHUB_API_TOKEN in header
	"fmt"
	"log"
	"os"
	"strconv"
)

// TODO: add fields, args for PATCH
// TODO: open text editor for text input

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
}

func main() {

	action := os.Args[1]
	if action == "GET" {
		owner, repo := os.Args[2], os.Args[3]
		// get all issues
		if len(os.Args) == 4 {
			res, err := github.GetIssues(owner, repo)
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(*res); i++ {
				issue := (*res)[i]
				printIssue(&issue)
			}
		}
		// get specific issue
		if len(os.Args) == 5 {
			issueNumber, _ := strconv.Atoi(os.Args[4])
			res, err := github.GetIssue(owner, repo, issueNumber)
			if err != nil {
				log.Fatal(err)
			}
			printIssue(res)
		}
	}

	// TODO: open editor to fill in body of Issue
	if action == "CREATE" {
		owner, repo := os.Args[2], os.Args[3]
		newIssue := &github.NewIssue{"Chapter 4 Post Test 2", "This is a another test"}
		res, err := github.CreateIssue(owner, repo, newIssue)
		if err != nil {
			log.Fatal(err)
		}
		printIssue(res)
	}

	if action == "CLOSE" {
		owner, repo := os.Args[2], os.Args[3]
		number, _ := strconv.Atoi(os.Args[4])
		res, err := github.CloseIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
		printIssue(res)
	}

	if action == "UPDATE" {
		owner, repo := os.Args[2], os.Args[3]
		number, _ := strconv.Atoi(os.Args[4])
		res, err := github.UpdateIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
		printIssue(res)
	}

}

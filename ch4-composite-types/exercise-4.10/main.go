package main

import (
	"ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	ageBuckets := make(map[string][]*github.Issue)
	now := time.Now()

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, issue := range result.Items {
		if now.Sub(issue.CreatedAt).Hours() < 24*30 {
			ageBuckets["< 1 Month"] = append(ageBuckets["< 1 Month"], issue)
		} else if now.Sub(issue.CreatedAt).Hours() < 24*365 {
			ageBuckets["< 1 Year"] = append(ageBuckets["< 1 Year"], issue)
		} else {
			ageBuckets["Older"] = append(ageBuckets["> 1 Year"], issue)
		}
	}
	for age, issues := range ageBuckets {
		fmt.Printf("Age: %v\n", age)
		for _, issue := range issues {
			fmt.Printf("#%-5d %9.9s %.55s %d\n", issue.Number, issue.User.Login, issue.Title)
		}
	}

}

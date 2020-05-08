package github

import "time"

const (
	ReposURL        = "https://api.github.com/repos"
	IssuesURL       = "https://api.github.com/issues"
	SearchIssuesURL = "https://api.github.com/search/issues"
	APITokenEnvVar  = "GITHUB_API_TOKEN"
)

type Issue struct {
	Number      int
	HTMLURL     string `json:"html_url"`
	Title       string
	Labels      []*Label
	State       string
	User        *User
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Body        string       // in Markdown format
	PullRequest *PullRequest `json:"pull_request"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Label struct {
	Id  int
	URL string
}

type PullRequest struct {
	URL     string
	HTMLURL string `json:"html_url"`
}

// TODO: the 3 below can be combined
type IssueState struct {
	State string `json:"state"`
}

type NewIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	// Labels    []string
	// Assignees []string
}

type IssueUpdate struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

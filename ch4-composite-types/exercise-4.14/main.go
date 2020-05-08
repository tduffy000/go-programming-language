package main

import (
	"ch4/github"
	"html/template"
	"log"
	"os"
)

var issueList = template.Must(template.New("issuelist").Parse(`
  <h1>Issues</h1>
  <table>
  <tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>User</th>
    <th>Title</th>
  </tr>
  {{range .}}
  <tr>
    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
    <td>{{.State}}</td>
    <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
    <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  </tr>
  {{end}}
  </table>
`))

func main() {
	owner, repo := os.Args[1], os.Args[2]
	res, err := github.GetIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}

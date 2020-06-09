package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// TODO: add more...
// TODO: handle duration
// TODO: make buttons for column names
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Soveig", "Smash", 2011, length("4m24s")},
}

var trackList = template.Must(template.New("tracklist").Parse(`
	<h1>Tracks Table</h1>
	<table>
	<tr style='text-align: left'>
		<th>Title</th>
		<th>Artist</th>
		<th>Album</th>
		<th>Year</th>
		<th>Length</th>
	</tr>
	{{ range . }}
	<tr>
		<td>{{ .Title }}</td>
		<td>{{ .Artist }}</td>
		<td>{{ .Album }}</td>
		<td>{{ .Year }}</td>
		<td>{{ .Length }}</td>
	</tr>
	{{ end }}
	</table>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type statefulSort struct {
	t         []*Track
	sortOrder []string
}

// this is SUPER hacky (but I couldn't figure out how to get around it)
func (x statefulSort) less(a, b *Track) bool {
	for _, fieldName := range x.sortOrder {
		ra, rb := reflect.ValueOf(a), reflect.ValueOf(b)
		valA, valB := reflect.Indirect(ra).FieldByName(fieldName), reflect.Indirect(rb).FieldByName(fieldName)
		if fieldName == "Title" || fieldName == "Artist" || fieldName == "Album" { // hard-coded :(
			if valA.String() != valB.String() {
				return valA.String() < valB.String()
			}
		} else { // how to deal with time.Duration?
			if valA.Int() != valB.Int() {
				return valA.Int() < valB.Int()
			}
		}
	}
	return false
}

func (x statefulSort) Len() int           { return len(x.t) }
func (x statefulSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x statefulSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }

func handler(w http.ResponseWriter, r *http.Request) {
	sortFields := r.URL.Query()["sort"]
	if len(sortFields) > 0 {
		sort.Sort(statefulSort{tracks, sortFields})
	}
	if err := trackList.Execute(w, tracks); err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
	}
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

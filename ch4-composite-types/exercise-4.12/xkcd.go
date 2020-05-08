package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	XkcdURL  = "https://xkcd.com"
	MaxComic = 400
)

type ComicDescription struct {
	Month       string
	Number      int `json:"num"`
	Link        string
	Year        string
	News        string
	SafeTitle   string `json:"safe_title"`
	Transcript  string
	Alternative string `json:"alt"`
	Image       string `json:"img"`
	Title       string
	Day         string
}

func GetComicDescription(number int) (*ComicDescription, error) {

	idx := strconv.Itoa(number)
	url := XkcdURL + "/" + idx + "/info.0.json"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get description failed: %s\n", resp.Status)
	}

	var description ComicDescription
	if err := json.NewDecoder(resp.Body).Decode(&description); err != nil {
		return nil, err
	}
	return &description, nil
}

func main() {

	outFile := "descriptions.json"

	if os.Args[1] == "BUILD" {
		var descriptions []*ComicDescription

		// get all the descriptions
		for i := 1; i <= MaxComic; i++ {
			desc, _ := GetComicDescription(i)
			fmt.Printf("Downloaded Comic number: %v of %v\n", desc.Number, MaxComic)
			descriptions = append(descriptions, desc)
		}
		// now write them out line by line
		os.Remove(outFile)
		os.Create(outFile)
		f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		for i, s := range descriptions {
			fmt.Printf("Writing out Comic number: %v\n", i)
			json, _ := json.Marshal(s)
			io.WriteString(f, string(json)+"\n")
		}

	} else if os.Args[1] == "SEARCH" {
		f, err := os.Open(outFile)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		line := 1
		for scanner.Scan() {
			searchArg, _ := strconv.Atoi(os.Args[2])
			if line == searchArg {
				fmt.Printf("Here's comic number %v: %v\n", line, scanner.Text())
			}
			line++
		}
	} else {
		log.Fatal("Only BUILD and SEARCH operations permitted.")
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type PackageInfo struct {
	Dir        string   `json:"Dir"`
	ImportPath string   `json:"ImportPath"`
	Name       string   `json:"Name"`
	GoFiles    []string `json:"GoFiles"`
	Imports    []string `json:"Imports,omitempty"`
	Deps       []string `json:"Deps,omitempty"`
}

func parse(b []byte) (*PackageInfo, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, b); err != nil {
		return nil, err
	}
	var info PackageInfo
	err := json.Unmarshal(buf.Bytes(), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func list(patterns []string) ([]*PackageInfo, error) {
	baseArgs := []string{"list", "-json"}
	var pkgs []*PackageInfo
	for _, pattern := range patterns {
		args := append(baseArgs, pattern)
		out, err := exec.Command("go", args...).Output()
		if err != nil {
			return nil, err
		}
		info, err := parse(out)
		if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, info)
	}
	return pkgs, nil
}

func main() {

	var patterns []string
	if len(os.Args[1:]) == 0 {
		patterns = append(patterns, "")
	} else {
		patterns = append(patterns, os.Args[1:]...)
	}
	// add cmd line args for searching
	looking := make(map[string]bool)
	for _, p := range patterns {
		looking[p] = true
	}

	// list the current workspace imports
	infos, err := list(patterns)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	}
	imports := infos[0].Imports
	fmt.Printf("Current workspace imports: %v\n", imports)

	// now list all transitive dependencies
	transitiveDeps := make(map[string]bool)
	for _, pkg := range imports {
		info, _ := list([]string{pkg})
		for _, d := range info[0].Deps {
			if looking[d] {
				transitiveDeps[pkg] = true
			}
		}
	}
	for imp, _ := range transitiveDeps {
		fmt.Printf("Import: %s has a dependency!\n", imp)
	}
}

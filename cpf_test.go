package main

import "strings"
import "testing"

func TestReadAndSavePaths(t *testing.T) {
	savePaths(".", []string{"a.go", "b.go"})
	paths := readPaths(".")
	if strings.Join(paths, " ") != "a.go b.go" {
		t.Fatal()
	}
}

package main

import "path/filepath"
import "io/ioutil"
import "strings"
import "fmt"
import "io"
import "os"

func main() {
	home := os.Getenv("HOME")
	args := os.Args[1:]
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	if len(args) == 0 {
		// put file
		paths := readPaths(home)
		copyFiles(wd, paths)
	} else {
		// save file paths
		var paths []string
		for _, arg := range args {
			if strings.Contains(arg, "*") {
				f, _ := filepath.Glob(arg)
				paths = append(paths, f...)
			} else {
				paths = append(paths, arg)
			}
		}

		// todo: remove repeat
		fmt.Println(paths)

		for i, v := range paths {
			paths[i] = filepath.Join(wd, v)
		}

		savePaths(home, paths)
	}
}

func savePaths(dir string, paths []string) {
	data := []byte(strings.Join(paths, "\n"))
	file := filepath.Join(dir, ".cpf_tmp_fps")
	ioutil.WriteFile(file, data, 0666)
}

func readPaths(dir string) []string {
	file := filepath.Join(dir, ".cpf_tmp_fps")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	return strings.Split(string(data), "\n")
}

func copyFiles(wd string, paths []string) {
	for _, path := range paths {
		filename := filepath.Join(wd, filepath.Base(path))
		copyFile(path, filename)
	}
}

func copyFile(from, dest string) {
	src, err := os.Open(from)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer src.Close()

	dst, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

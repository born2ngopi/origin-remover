package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	colorReset := "\033[0m"

	colorYellow := "\033[33m"

	var paths []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			// continue if directory
			return nil
		}

		// check file extension if .orig
		if filepath.Ext(path) == ".orig" {
			// append to paths
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
Files:
	if len(paths) == 0 {
		fmt.Println("No files found")
		os.Exit(0)
	}

	fmt.Printf("Found %d files on : \n", len(paths))
	for _, path := range paths {
		fmt.Println(string(colorYellow), path, string(colorReset))
	}

	fmt.Printf("\n\n")
	fmt.Println("(d) Deleted, (u) update list / remove from list, (q) quit")
	var input string
	fmt.Scanln(&input)

	if input == "d" {
		for _, path := range paths {
			os.Remove(path)
		}
		fmt.Println("Deleted")
	} else if input == "u" {
		fmt.Println("copy path you wont to remove from list")
		var inputPath string
		fmt.Scanln(&inputPath)
		for i, path := range paths {
			if path == inputPath {
				paths = append(paths[:i], paths[i+1:]...)
			}
		}
		fmt.Println("removed")
		fmt.Printf("\n\n")
		goto Files
	} else if input == "q" {
		fmt.Println("Quit")
		os.Exit(0)
	}
}

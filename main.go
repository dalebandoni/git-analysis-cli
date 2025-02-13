package main

import (
	"flag"
	"fmt"
)

func scan(path string) {
	fmt.Println("scan")
}

func stats(email string) {
	fmt.Println("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "add", "", "add a new folder to scan for Git repositories")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}

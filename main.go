package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	GhUser     string = "GITHUB_USER"
	GhToken    string = "GITHUB_TOKEN"
	UrlPattern string = "https://%s:%s@github.com/%s/%s.git"
)

func getRepo(args []string) (string, error) {
	v := args[0]

	var e error
	if strings.TrimSpace(v) == "" {
		e = fmt.Errorf("invalid repo-name")
	}

	return v, e
}

func getInfo() (string, string, error) {
	u := os.Getenv(GhUser)
	t := os.Getenv(GhToken)

	var e error
	if u == "" || t == "" {
		e = fmt.Errorf("$GITHUB_USER is empty or $GITHUB_TOKEN is empty or both empty")
	}

	return u, t, e
}

func makeUrl(user, token, repo string) (string, error) {
	return fmt.Sprintf(UrlPattern, user, token, user, repo), nil
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("no repo-name specified")
	}

	repo, err := getRepo(args)
	if err != nil {
		log.Fatal(err)
	}

	user, token, err := getInfo()
	if err != nil {
		log.Fatal(err)
	}

	url, err := makeUrl(user, token, repo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
}

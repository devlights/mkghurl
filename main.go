package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	ghUser     string = "GITHUB_USER"
	ghToken    string = "GITHUB_TOKEN"
	urlPattern string = "https://%s:%s@github.com/%s/%s.git"
	version    string = "v1.0.3"
)

var (
	showVersion = flag.Bool("v", false, "show version")
)

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: mkghurl repo-name [flags]\n")
	flag.PrintDefaults()
}

func errorExit(e error) {
	_, _ = fmt.Fprintf(os.Stderr, "%s", e)
	os.Exit(1)
}

func getRepo(args []string) (string, error) {
	v := args[0]

	var e error
	if strings.TrimSpace(v) == "" {
		e = fmt.Errorf("invalid repo-name")
	}

	return v, e
}

func getInfo() (string, string, error) {
	u := os.Getenv(ghUser)
	t := os.Getenv(ghToken)

	var e error
	if u == "" || t == "" {
		e = fmt.Errorf("$GITHUB_USER is empty or $GITHUB_TOKEN is empty or both empty")
	}

	return u, t, e
}

func makeUrl(user, token, repo string) (string, error) {
	return fmt.Sprintf(urlPattern, user, token, user, repo), nil
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", version)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		usage()
		os.Exit(1)
	}

	repo, err := getRepo(flag.Args())
	if err != nil {
		errorExit(err)
	}

	user, token, err := getInfo()
	if err != nil {
		errorExit(err)
	}

	url, err := makeUrl(user, token, repo)
	if err != nil {
		errorExit(err)
	}

	_, _ = fmt.Fprintln(os.Stdout, url)
	os.Exit(0)
}

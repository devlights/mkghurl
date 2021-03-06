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
	version    string = "v1.0.7"
)

var (
	showVersion  = flag.Bool("version", false, "show version")
	noNewLine    = flag.Bool("n", false, "with no newline")
	withGitClone = flag.Bool("g", false, "with 'git clone' prefix")
)

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: mkghurl [options] repo-name\n")
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

func makeURL(user, token, repo string) (string, error) {
	return fmt.Sprintf(urlPattern, user, token, user, repo), nil
}

func withPrefix(url string, gitClone bool) (string, error) {
	var e error
	if !gitClone {
		return url, e
	}

	return fmt.Sprintf("git clone %s", url), e
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", version)
		os.Exit(0)
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

	url, err := makeURL(user, token, repo)
	if err != nil {
		errorExit(err)
	}

	urlWithPrefix, err := withPrefix(url, *withGitClone)
	if err != nil {
		errorExit(err)
	}

	if *noNewLine {
		_, _ = fmt.Fprint(os.Stdout, urlWithPrefix)
	} else {
		_, _ = fmt.Fprintln(os.Stdout, urlWithPrefix)
	}

	os.Exit(0)
}

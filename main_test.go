package main

import "testing"

func TestGetRepo(t *testing.T) {
	expected := "testrepo"
	got, err := getRepo([]string{expected})

	if err != nil {
		t.Error(err)
	}

	if expected != got {
		t.Errorf("expected:%s\tgot:%s", expected, got)
	}
}

func TestGetInfo(t *testing.T) {
	_, _, err := getInfo()

	if err != nil {
		t.Error(err)
	}
}

func TestMakeUrl(t *testing.T) {
	expected := "https://u:t@github.com/u/r.git"
	got, err := makeUrl("u", "t", "r")

	if err != nil {
		t.Error(err)
	}

	if expected != got {
		t.Errorf("expected:%s\tgot:%s", expected, got)
	}
}

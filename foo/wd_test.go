package main

import (
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium"
)

func TestStart(t *testing.T) {
	caps := selenium.Capabilities{}
	wd, err := webtest.NewWebDriverSession(caps)
	if err != nil {
		t.Fatalf("error starting webdriver: %s", err)
	}
	if err := wd.Quit(); err != nil {
		t.Errorf("Error quitting webdriver: %s", err)
	}
}

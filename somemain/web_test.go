package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium"
)

func TestOkay(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	caps := selenium.Capabilities{}

	wd, err := webtest.NewWebDriverSession(caps)
	if err != nil {
		t.Fatal(err)
	}
	err = wd.Get(srv.URL + "/")
	if err != nil {
		t.Fatalf("error while GETing /: %s", err)
	}
	if err := wd.Quit(); err != nil {
		t.Errorf("Error quiting webdriver: %s", err)
	}
}

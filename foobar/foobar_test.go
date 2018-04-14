package foobar

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium"
)

func TestOkay(t *testing.T) {
	caps := selenium.Capabilities{}
	wd, err := webtest.NewWebDriverSession(caps)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := wd.Quit(); err != nil {
			t.Errorf("Error quitting webdriver: %v", err)
		}
	}()

	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<!DOCTYPE html>\n<html><body><p>hi</p></body></html>"))
	}))

	// Comment this StartTLS and and uncomment Start to see this work.
	srv.StartTLS()
	// srv.Start()

	t.Logf("srv.URL: %s", srv.URL)
	wd.Get(srv.URL + "/")

	el, err := wd.FindElement(selenium.ByCSSSelector, "p")
	if err != nil {
		t.Fatalf("error finding p: %s", err)
	}
	txt, err := el.Text()
	if err != nil {
		t.Fatalf("error pulling text of <p>: %s", err)
	}
	if txt != "hi" {
		t.Errorf("txt: want %#v, got %#v", "hi", txt)
	}
}

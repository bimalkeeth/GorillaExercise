package main

import (
	"fmt"
	"github.com/gorilla/reverse"
	"net/http"
	"net/url"
)

func main() {
	y := &url.URL{
		Scheme:   "http",
		Host:     "localhost:9999",
		Path:     "/foo/42",
		RawQuery: "buz=42",
	}

	r := &http.Request{URL: y}
	p, _ := reverse.NewRegexpPath("/foo/[0-9]+")
	q := reverse.NewQuery(map[string]string{"buz": "43"})

	a := reverse.NewAll([]reverse.Matcher{p, q})

	fmt.Println("Match", a.Match(r))

}

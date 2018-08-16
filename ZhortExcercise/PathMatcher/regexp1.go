package main

import (
	"fmt"
	"github.com/gorilla/reverse"
	"net/url"
)

func main() {

	regexp, _ := reverse.CompileRegexp(`/foo/(d+)`)
	r, _ := regexp.Revert(url.Values{"": {"42"}})
	fmt.Println(r)
}

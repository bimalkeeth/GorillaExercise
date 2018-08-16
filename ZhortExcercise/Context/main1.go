package main

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

const key = "MyKey"

func main() {
	request1 := &http.Request{}
	request2 := &http.Request{}
	context.Set(request1, key, "foo")
	context.Set(request2, key, "bar")

	fmt.Println(context.Get(request1, key))
}

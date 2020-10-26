package main

import (
	"github.com/goandfootball/test-api/routers"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:8080", routers.Handler())
}
package routers

import (
	"fmt"
	"net/http"
)

func GetTest (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
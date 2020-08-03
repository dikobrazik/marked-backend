package main

import (
	"net/http"
)

func main() {
	r := Router()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

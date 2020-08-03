package main

import "net/http"

type auth struct {
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("Authorization")
		if auth != "" {

		}
	})
}

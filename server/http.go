package server

import (
	"net/http"
)

//
func Serve() error {
	server := &http.Server{
		Addr: ":8080",
	}

	return server.ListenAndServe()
}

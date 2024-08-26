package main

import (
	"fmt"
	"goapi/api/router"
	"net/http"
)

func main() {
	r := router.New()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", "8080"),
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server startup failure")
	}
}

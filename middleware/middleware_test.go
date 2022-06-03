package middleware

import (
	"fmt"
	"testing"
)

func MiddlewareTest(t *testing.T) {
	middlewareChain := MiddlewareChain{
		storage: map[string]MiddlewareStorage{
			"log":    logging,
			"filter": filterContentType,
		},
	}

	// Middleware(http.DefaultServeMux, "")

	fmt.Println(middlewareChain)
}

// func loggingTest(t *testing.T) {

// 	logging(http.DefaultServeMux)
// }

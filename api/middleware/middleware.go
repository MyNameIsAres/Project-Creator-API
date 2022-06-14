package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type MiddlewareStorage func(http.Handler) http.Handler

type MiddlewareChain struct {
	storage map[string]MiddlewareStorage
}

type MiddlewareStorageGroup struct {
	storage map[string][]MiddlewareStorage
}

//? Do we need this? And do we need this here?
func CreateMiddlewareChain() *MiddlewareChain {
	return &MiddlewareChain{
		storage: map[string]MiddlewareStorage{
			"log":    logging,
			"filter": filterContentType,
			"foobar": loggingMiddleware,
		},
	}
}

//! This style has to be reviewed. It may not be appropriate
//! Not up to standards with Go.. or just unecessarily complicated.
func CreateMiddlewareGroups() *MiddlewareStorageGroup {
	return &MiddlewareStorageGroup{
		storage: map[string][]MiddlewareStorage{
			"api": {
				logging,
				filterContentType,
			},
		},
	}
}

func validateMiddlewareKeys(keys ...string) []MiddlewareStorage {
	middlewareChain := CreateMiddlewareChain()
	var validatedKeys []MiddlewareStorage

	//? Best case scenario is O(1) if there's 1 item to loop over, I think?
	for _, key := range keys {
		if value, exists := middlewareChain.storage[key]; exists {
			validatedKeys = append(validatedKeys, value)
		}
	}

	return validatedKeys
}

func validateGroupId(groupId string, middlewareGroup *MiddlewareStorageGroup) ([]MiddlewareStorage, error) {
	var validatedGroupMiddleware []MiddlewareStorage

	if value, exists := middlewareGroup.storage[groupId]; exists {
		validatedGroupMiddleware = append(validatedGroupMiddleware, value...)
	} else {
		return nil, errors.New("No middleware group with ID: " + groupId)
	}

	return validatedGroupMiddleware, nil
}

func Middleware(handler http.Handler, keys ...string) http.Handler {
	// TODO Maybe move.. because now constantly call it?
	// ? Do we call all middleware on each request? Do filter some out?
	validatedKeys := validateMiddlewareKeys(keys...)

	return runMiddleware(handler, validatedKeys)
}

// ! Do not allow multiple middleware groups for now.
func MiddlewareGroup(handler http.Handler, groupId string) http.Handler {
	middlewareGroup := CreateMiddlewareGroups()
	validatedGroupMiddleware, err := validateGroupId(groupId, middlewareGroup)
	if err != nil {
		panic(err)
	}
	return runMiddleware(handler, validatedGroupMiddleware)
}

func runMiddleware(handler http.Handler, middleware []MiddlewareStorage) http.Handler {
	for i := range middleware {
		handler = middleware[len(middleware)-1-i](handler)
	}
	return handler
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Barfoo")

		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Foobar")
		next.ServeHTTP(w, r)
	})

}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Currently in the filterContentType middleware")
		handler.ServeHTTP(w, r)
	})

}

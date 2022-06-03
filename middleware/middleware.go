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
//! Nor up to standards with Go.. or just unecessarily complicated.
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

func validateGroupId(groupId string) ([]MiddlewareStorage, error) {
	middlewareGroup := CreateMiddlewareGroups()
	var validatedGroupMiddleware []MiddlewareStorage

	if value, exists := middlewareGroup.storage[groupId]; exists {
		validatedGroupMiddleware = append(validatedGroupMiddleware, value...)
	} else {
		return nil, errors.New("No middleware group with ID: " + groupId)
	}

	return validatedGroupMiddleware, nil
}

func Middleware(handler http.Handler, keys ...string) http.Handler {

	fmt.Println("First handler:", handler)

	// TODO Maybe move.. because now constantly call it?
	// ? Do we call all middleware on each request? Do filter some out?
	validatedKeys := validateMiddlewareKeys(keys...)

	// Iterate over each key and handle the functions associated.
	foo := runMiddleware(handler, validatedKeys)

	return foo
}

// ! Do not allow multiple middleware groups for now.
func MiddlewareGroup(handler http.Handler, groupId string) http.Handler {

	validatedGroupMiddleware, err := validateGroupId(groupId)
	if err != nil {
		panic(err)
	}
	return runMiddleware(handler, validatedGroupMiddleware)
}

func runMiddleware(handler http.Handler, middleware []MiddlewareStorage) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handleFunc := range middleware {
			handler = handleFunc(handler)

			handler.ServeHTTP(w, r)
		}

		// next.ServeHTTP(w, r)
	})
	// var foo http.Handler

	// for _, i := range middleware {

	// 	handler = i(handler)

	// 	foo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 		handler.ServeHTTP(w, r)
	// 	})
	// }

	// // 	handler = middleware[len(middleware)-1-i](handler)

	// // 	fmt.Println(handler, " The handler in question")
	// // }

	// fmt.Println("Is this still original handler", handler)

	// return foo
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HLOOH")

		// next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {

	x := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Foobar")

		// next.ServeHTTP(w, r)
	})

	return x

}

// middleware to check the content type header
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Currently in the filterContentType middleware")
		// // fmt.Println(r.Header)
		// if r.Header.Get("Content-Type") != "application/json" {
		// 	w.WriteHeader(http.StatusUnsupportedMediaType)
		// 	w.Write([]byte("415 - Header Content-type missing"))
		// 	return
		// }

		// fmt.Println(w.Header())
		// fmt.Println(r.Response)

		// handler.ServeHTTP(w, r)
	})

}

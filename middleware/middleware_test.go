package middleware

import (
	"net/http"
	"testing"
)

//! Test functions are incomplete
// TODO Continue to create middleware tests as the project progresses.

func TestValidateGroupIdReturnsError(t *testing.T) {
	middlewareStorage := MiddlewareStorageGroup{
		storage: map[string][]MiddlewareStorage{
			"api": {
				func(handler http.Handler) http.Handler {
					return handler
				},
			},
		},
	}

	keys, groupIdError := validateGroupId("hello", &middlewareStorage)

	if keys != nil {
		t.Errorf("Wanted error, got %d number of keys", len(keys))
	}

	if groupIdError == nil {
		t.Errorf("Expected error, got nil instead")
	}
}

func TestValidateGroupId(t *testing.T) {
	middlewareStorage := MiddlewareStorageGroup{
		storage: map[string][]MiddlewareStorage{
			"api": {
				func(handler http.Handler) http.Handler {
					return handler
				},
			},
		},
	}

	keys, groupIdError := validateGroupId("api", &middlewareStorage)

	if keys == nil {
		t.Errorf("Wanted keys got nil instead.")
	}

	if groupIdError != nil {
		t.Errorf("Expected keys got error instead")
	}
}

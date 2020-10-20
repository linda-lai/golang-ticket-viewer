package mocks

import "net/http"

var (
	// The GetDoFunc can hold any value that is a function taking in
	// an argument of a pointer to an http.Request and return either
	// a pointer to an http.Response or an error.
	// 	GetDoFunc fetches the mock client's `Do` func
	GetDo func(req *http.Request) (*http.Response, error)
)

// MockClient struct is the mock client
type MockClient struct {
	GetDoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDo(req)
}

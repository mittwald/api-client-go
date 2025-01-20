package httpclient

import "net/http"

var _ RequestRunner = &http.Client{}

// RequestRunner is an abstraction for HTTP clients. The usual *http.Client
// type natively implements this interface; however, this interface also allows
// you to swap it client with your own implementation, if required.
type RequestRunner interface {
	Do(*http.Request) (*http.Response, error)
}

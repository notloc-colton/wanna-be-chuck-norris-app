// This is a stub package that takes the place of a custom made http solution
package httpclient

import "net/http"

type Client interface {
	Get(url string) (resp *http.Response, err error)
}

func StandardClient() *http.Client {
	return &http.Client{}
}

package testhttpclient

import "math/rand"

type TestHTTPClient struct{}

func (client TestHTTPClient) Get(url string) (any, error) {
	return rand.Float64(), nil
}

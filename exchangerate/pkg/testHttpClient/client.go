package testhttpclient

import "math/rand"

type TestHTTPClient struct{}

func (client TestHTTPClient) Get(url string) (any, error) {
	rate := rand.Float32() * 100
	return rate, nil
}

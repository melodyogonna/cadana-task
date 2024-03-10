package testhttpclient

type TestHTTPClient struct{}

func (client TestHTTPClient) Get(url string) (any, error) {
	return 1.0, nil
}

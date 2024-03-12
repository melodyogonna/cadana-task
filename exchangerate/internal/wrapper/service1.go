package wrapper

import (
	"os"
)

type serviceOne struct {
	httpClient HttpClient
	secretkey  string
	url        string
}

func (s1 serviceOne) GetExchangeRate(pair string) <-chan result {
	c := make(chan result)

	go func() {
		res, err := s1.httpClient.Get(s1.url + "/?currency-pair=" + pair)
		if err != nil {
			c <- result{Err: err, Value: 0}
		}
		c <- result{Err: nil, Value: res.(float64)}
	}()

	return c
}

func initService1(httpClient HttpClient) serviceOne {
	url := "https://api.serviceOne.com"
	secretkey := os.Getenv("SERVICE_SECRET")
	return serviceOne{httpClient, secretkey, url}
}

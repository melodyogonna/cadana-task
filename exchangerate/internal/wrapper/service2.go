package wrapper

import (
	"os"
)

type serviceTwo struct {
	httpClient HttpClient
	secretkey  string
	url        string
}

func (s2 serviceTwo) GetExchangeRate(pair string) <-chan result {
	c := make(chan result)

	go func() {
		res, err := s2.httpClient.Get(s2.url + "/?pair=" + pair)
		if err != nil {
			c <- result{Err: err, Value: 0}
		}

		c <- result{Err: err, Value: res.(float64)}

	}()

	return c
}

func initService2(httpClient HttpClient) serviceTwo {
	url := "https://api.serviceOne.com"
	secretkey := os.Getenv("SERVICE_SECRET")
	return serviceTwo{httpClient, secretkey, url}
}

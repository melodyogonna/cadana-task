package wrapper

import (
	"os"
)

type serviceOne struct {
	httpClient HttpClient
	secretkey  string
	url        string
}

func (s1 serviceOne) GetExchangeRate(pair string, c chan float32, e chan error) {
	res, err := s1.httpClient.Get(s1.url + "/?currency-pair=" + pair)
	if err != nil {
		e <- err
	}

	c <- float32(res.(float64))
}

func initService1(httpClient HttpClient) serviceOne {
	url := "https://api.serviceOne.com"
	secretkey := os.Getenv("SERVICE_SECRET")
	return serviceOne{httpClient, secretkey, url}
}

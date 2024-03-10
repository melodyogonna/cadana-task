package wrapper

import (
	"os"
)

type serviceTwo struct {
	httpClient HttpClient
	secretkey  string
	url        string
}

func (s2 serviceTwo) GetExchangeRate(pair string, c chan float32, e chan error) {
	res, err := s2.httpClient.Get(s2.url + "/?pair=" + pair)
	if err != nil {
		e <- err
	}

	c <- float32(res.(float64))
}

func initService2(httpClient HttpClient) serviceTwo {
	url := "https://api.serviceOne.com"
	secretkey := os.Getenv("SERVICE_SECRET")
	return serviceTwo{httpClient, secretkey, url}
}

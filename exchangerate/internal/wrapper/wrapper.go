package wrapper

type Wrapper interface {
	GetExchangeRate(pair string, c chan float32, e chan error)
}

type HttpClient interface {
	Get(url string) (any, error)
}

func Init(client HttpClient) []Wrapper {
	var wrappers []Wrapper

	wrappers = append(wrappers, initService1(client))
	wrappers = append(wrappers, initService2(client))

	return wrappers
}

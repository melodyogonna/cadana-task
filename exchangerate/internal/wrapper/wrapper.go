package wrapper

type result struct {
	Err   error
	Value float32
}

// Wraps external services that we can query for exchange rate.
// returns a readonly channel where the result of the retrieve can get passed.
type Wrapper interface {
	GetExchangeRate(pair string) <-chan result
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

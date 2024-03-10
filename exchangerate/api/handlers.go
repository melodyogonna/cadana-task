package api

import (
	"encoding/json"
	"errors"
	"exchangerate/internal/wrapper"
	testhttpclient "exchangerate/pkg/testHttpClient"
	"net/http"
)

func getExchangeRate(pair string) (float32, error) {
	// get exchange rate for pair
	httpClient := testhttpclient.TestHTTPClient{}
	services := wrapper.Init(httpClient)

	result := make(chan float32)
	err := make(chan error)
	for _, service := range services {
		go service.GetExchangeRate(pair, result, err)
	}

	var outPut float32
	var numErrors int

	for {
		select {
		case r := <-result:
			outPut = r
			break

		case <-err:
			if numErrors == 0 {
				numErrors += 1
				continue
			} else {
				numErrors += 1
				break
			}
		}
	}

	if numErrors > 1 {
		return 0, errors.New("Unable to retrieve exchange rate")
	}

	return outPut, nil
}

type currencyPair struct {
	Pair string `json:"currency-pair"`
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	body, err := readBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parsed := currencyPair{}
	if err = toJson(body, &parsed); err != nil {
		http.Error(w, "Not valid json", http.StatusBadRequest)
		return
	}

	price, err := getExchangeRate(parsed.Pair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	var res map[string]float32

	res[parsed.Pair] = price

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(json))

}

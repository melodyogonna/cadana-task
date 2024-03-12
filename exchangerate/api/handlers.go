package api

import (
	"context"
	"encoding/json"
	"exchangerate/internal/wrapper"
	testhttpclient "exchangerate/pkg/testHttpClient"
	"log"
	"net/http"
)

func launchServices(ctx context.Context, pair string) (float32, error) {
	// get exchange rate for pair
	httpClient := testhttpclient.TestHTTPClient{}
	services := wrapper.Init(httpClient)
	resultChan := make(chan float32)

	for _, service := range services {
		go func(s wrapper.Wrapper) {
			select {
			case <-ctx.Done():
				break
			case result := <-s.GetExchangeRate(pair):
				if result.Err != nil {
					log.Print(result.Err)
					break
				}
				resultChan <- result.Value

			}
		}(service)
	}

	ouput := <-resultChan

	return ouput, nil

}

func getExchangeRate(pair string) (float32, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := launchServices(ctx, pair)

	if err != nil {
		return 0, err
	}

	return result, nil
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

	if parsed.Pair == "" {
		message := map[string]string{"message": "'currency-pair' not passed"}
		m, _ := json.Marshal(message)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(m)
		return
	}

	price, err := getExchangeRate(parsed.Pair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	res := make(map[string]float32)

	res[parsed.Pair] = price

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(json))

}

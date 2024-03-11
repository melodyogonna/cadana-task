package accounting

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func retrieveRate(pair string) (float64, error) {
	body := map[string]string{"currency-pair": pair}
	marsharledJson, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", "", bytes.NewReader(marsharledJson))
	if err != nil {
		return 0, errors.New("Error while creating request")
	}

	req.Header.Set("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("")
	}

	var res map[string]float64
	if err := json.Unmarshal(r, &res); err != nil {
		return 0, err
	}

	return res[pair], nil

}

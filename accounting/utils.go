package accounting

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
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
	if err != nil {
		return 0, err
	}

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

func getJsonFromFileSystem(jsonPath string) (*[]Person, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var p []Person

	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func InitPersonsFromJson(jsonPath string) (Persons, error) {
	persons, err := getJsonFromFileSystem(jsonPath)
	if err != nil {
		return Persons{}, err
	}

	return Persons{Data: *persons}, nil

}

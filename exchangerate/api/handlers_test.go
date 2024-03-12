package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExchangeRateReturned(t *testing.T) {
	body := bytes.NewReader([]byte(`{"currency-pair":"USD-GBP"}`))
	req := httptest.NewRequest("POST", "/exchange-rate", body)
	res := httptest.NewRecorder()

	requestHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Error: '%d' != '%d'", res.Code, http.StatusOK)
	}

	if res.Header().Get("content-type") != "application/json" {
		t.Errorf("Bad content-type returned.")
	}
}

func TestErrorOnInvalidJson(t *testing.T) {
	body := bytes.NewReader([]byte(`{"currency-pair":"USD-GBP}`))
	req := httptest.NewRequest("POST", "/exchange-rate", body)
	res := httptest.NewRecorder()

	requestHandler(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("Error: '%d' != '%d'", res.Code, http.StatusBadRequest)
	}
}

func TestErrorOnInvalidPayload(t *testing.T) {
	body := bytes.NewReader([]byte(`{"currency":"USD-GBP"}`))
	req := httptest.NewRequest("POST", "/exchange-rate", body)
	res := httptest.NewRecorder()

	requestHandler(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("Error: '%d' != '%d'", res.Code, http.StatusBadRequest)
	}
}

package cryptomus

import "net/http"

type Cryptomus struct {
	merchantId string
	apiKey     string
	httpClient *http.Client
}

func New(httpClient *http.Client, merchantId string, apiKey string) *Cryptomus {
	return &Cryptomus{
		merchantId: merchantId,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

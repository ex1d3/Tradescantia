package cryptomus

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const BASE_API_URL = "https://api.cryptomus.com/v1"

func (c *Cryptomus) call(method string, endpoint string, data any) (*http.Response, error) {
	body, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	sign := c.sign(body)

	req, err := http.NewRequest(method, BASE_API_URL+endpoint, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("merchant", c.merchantId)
	req.Header.Set("sign", sign)

	return c.httpClient.Do(req)
}

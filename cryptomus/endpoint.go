package cryptomus

import (
	"io"
	"log"
)

const testWebhookEndpoint = "/test-webhook/payment"

type TestWebhookDto struct {
	UrlCallback string `json:"url_callback"`
	Currency    string `json:"currency"`
	Network     string `json:"network"`
	UUID        string `json:"uuid"`
	Status      string `json:"status"`
}

func (c *Cryptomus) TestWebhook(dto TestWebhookDto) {
	res, err := c.call("POST", testWebhookEndpoint, dto)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		log.Fatalln(string(body))
	}

	log.Printf("Test webhook successfully sent on %s\n", dto.UrlCallback)
}

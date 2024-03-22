package main

import (
	"log"
	"net/http"
	"tradescantia/cli"
	"tradescantia/config"
	"tradescantia/cryptomus"
)

func main() {
	cfg, err := config.Parse()

	if err != nil {
		log.Fatalln(err)
	}

	cliArgs := cli.Parse()
	cliArgs.Validate()

	webhookDto := cryptomus.TestWebhookDto{
		UrlCallback: cfg.CallbackUrl,
		Currency:    cliArgs.Currency,
		Network:     cliArgs.Network,
		UUID:        cliArgs.InvoiceUUID,
		Status:      cliArgs.Status,
	}

	httpClient := http.Client{}

	cryptomusClient := cryptomus.New(&httpClient, cfg.MerchantId, cfg.ApiKey)
	cryptomusClient.TestWebhook(webhookDto)
}

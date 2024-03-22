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

	cliArgs := cli.ParseArgs()
	cliArgs.Validate()

	webhookDto := cryptomus.TestWebhookDto{
		UrlCallback: cfg.CallbackUrl,
		Currency:    "USDT",
		Network:     "TRON",
		UUID:        cliArgs.InvoiceUUID,
		Status:      "paid",
	}

	httpClient := http.Client{}

	cryptomusClient := cryptomus.New(&httpClient, cfg.MerchantId, cfg.ApiKey)
	cryptomusClient.TestWebhook(webhookDto)
}

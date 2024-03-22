package cli

import "flag"

func ParseArgs() Args {
	invoiceUUID := flag.String("invoiceUUID", "", "Invoice UUID in cryptomus system")
	callbackUrl := flag.String("callbackUrl", "", "Webhook callback url")

	flag.Parse()

	return Args{
		InvoiceUUID: *invoiceUUID,
		CallbackUrl: *callbackUrl,
	}
}

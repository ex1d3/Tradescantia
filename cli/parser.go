package cli

import "flag"

func Parse() *Args {
	currency := flag.String("currency", "", "Currency of invoice")
	network := flag.String("network", "", "Network of currency")
	status := flag.String("status", "", "Invoice status")
	invoiceUUID := flag.String("invoiceUUID", "", "Invoice UUID in cryptomus system")

	flag.Parse()

	return &Args{
		InvoiceUUID: *invoiceUUID,
		Currency:    *currency,
		Network:     *network,
		Status:      *status,
	}
}

package cli

import "flag"

func ParseArgs() Args {
	invoiceUUID := flag.String("invoiceUUID", "", "Invoice UUID in cryptomus system")

	flag.Parse()

	return Args{
		InvoiceUUID: *invoiceUUID,
	}
}

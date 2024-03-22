package cli

import (
	"fmt"
	"log"
	"reflect"
)

type Args struct {
	InvoiceUUID string `json:"invoiceUUID"`
}

// Validate Iterates over struct, and throws error
// if any field of struct is ""
func (args Args) Validate() {
	value := reflect.ValueOf(args)

	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Interface() == "" {
			arg := value.Type().Field(i).Tag.Get("json")

			log.Fatalln(fmt.Sprintf("\"%s\" arg is not passed", arg))
		}
	}
}

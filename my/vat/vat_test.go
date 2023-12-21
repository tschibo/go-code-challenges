package vat

import (
	"fmt"
	"testing"
)

func TestVat1(t *testing.T) {
	price := 100.00
	vatRate := 0.19
	fmt.Println(RenderInvoice(price, vatRate))
	//Output:
	//Netto: 100.00 €
	//VAT (0.19): 19.00 €
	//Total: 119.00 €
}

func TestVat2(t *testing.T) {
	price := 100.00
	vatRate := 0.19
	fmt.Println(RenderInvoice(price, vatRate))
	//Output:
	//Netto: 87.00 €
	//VAT (0.19): 16.53 €
	//Total: 103.53 €
}

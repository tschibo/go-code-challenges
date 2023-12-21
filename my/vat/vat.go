package vat

import (
	"fmt"
	"math"
)

func RenderInvoice(price, vatRate float64) string {
	vat := math.Round(price*vatRate*100) / 100
	gross := price + vat
	invoice := fmt.Sprintf("Netto: %.2f €\nVAT (%.2f): %.2f €\nTotal: %.2f", price, vatRate, vat, gross)
	return invoice
}

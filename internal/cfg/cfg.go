package cfg

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/raphael-foliveira/invoice-generator/internal/crawler"
)

func ReadInvoiceData() (*crawler.InvoiceData, error) {
	fmt.Println("Reading invoice data from config.json")
	b, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Invoice data must be provided in config.json")
		return nil, fmt.Errorf("cfg.ReadInvoiceData: %w", err)
	}
	invoiceData := new(crawler.InvoiceData)
	return invoiceData, json.Unmarshal(b, invoiceData)
}

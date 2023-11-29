package crawler

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type InvoiceData struct {
	InvoiceCode        int    `json:"invoiceCode"`
	Provider           string `json:"provider"`
	CNPJ               string `json:"cnpj"`
	Address            string `json:"address"`
	Email              string `json:"email"`
	Client             string `json:"client"`
	ClientCNPJ         string `json:"clientCnpj"`
	Currency           string `json:"currency"`
	Amount             string `json:"amount"`
	DueDate            string `json:"dueDate"`
	ServiceTitle       string `json:"serviceTitle"`
	ServiceDescription string `json:"serviceDescription"`
}

var invoiceUrl = "https://invoice.agilize.com.br/"

func readInvoiceData() (*InvoiceData, error) {
	b, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Invoice data must be provided in config.json")
		return nil, err
	}
	invoiceData := new(InvoiceData)
	return invoiceData, json.Unmarshal(b, invoiceData)
}

func getInvoicePage() *rod.Page {
	browser := rod.New().MustConnect().MustPage("")
	downloadPath := os.Args[1]
	_ = proto.PageSetDownloadBehavior{
		Behavior:     proto.PageSetDownloadBehaviorBehaviorAllow,
		DownloadPath: downloadPath,
	}.Call(browser)
	return browser.Browser().MustPage(invoiceUrl)
}

func fillInvoiceData(page *rod.Page, invoiceData *InvoiceData) {
	page.MustElement(`[name="invoiceCode"]`).MustInput(strconv.Itoa(invoiceData.InvoiceCode))
	page.MustElement("#companieName").MustInput(invoiceData.Provider)
	page.MustElement(`[data-testid="cnpj"]`).MustInput(invoiceData.CNPJ)
	page.MustElement("#address").MustInput(invoiceData.Address)
	page.MustElement("#client-email").MustInput(invoiceData.Email)
	page.MustElement("#customerName").MustInput(invoiceData.Client)
	page.MustElement(`[name="customerCnpjCpf"]`).MustInput(invoiceData.ClientCNPJ)
	page.MustElement(`button[for="opening-balance"]`).MustClick()
	currencyList := page.MustElement("#prefix-dropdown")
	currencies := currencyList.MustElements("li")
	for _, currency := range currencies {
		if currency.MustText() == invoiceData.Currency {
			currency.MustClick()
			break
		}
	}
	page.MustElement("#opening-balance").MustInput(invoiceData.Amount)
	page.MustElement("#fatura-vencimento").MustInput(invoiceData.DueDate)
	page.MustElement(`[name="serviceTitle"]`).MustInput(invoiceData.ServiceTitle)
	page.MustElement("#description").MustInput(invoiceData.ServiceDescription)
}

func clickDownload(page *rod.Page) {
	pageButtons := page.MustElements(`button`)
	for _, pageButton := range pageButtons {
		if pageButton.MustText() == "Baixar invoice" {
			pageButton.MustClick()
			break
		}
	}
}

func Run() error {
	fmt.Println("Reading invoice data from config.json")
	invoiceData, err := readInvoiceData()
	if err != nil {
		return err
	}
	fmt.Println("Starting crawler")
	page := getInvoicePage()
	fmt.Println("Filling invoice form")
	fillInvoiceData(page, invoiceData)
	fmt.Println("Downloading invoice")
	clickDownload(page)
	time.Sleep(5 * time.Second)
	fmt.Println("Closing")
	return page.Close()
}

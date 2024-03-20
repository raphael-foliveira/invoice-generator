package crawler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const invoiceUrl = "https://invoice.agilize.com.br/"

type crawler struct {
	invoiceData  InvoiceData
	downloadPath string
	*rod.Page
}

func New(invoiceData *InvoiceData, downloadPath string) *crawler {
	c := &crawler{
		invoiceData:  *invoiceData,
		downloadPath: downloadPath,
	}
	c.mustGetInvoicePage()
	return c
}

func (c *crawler) mustGetInvoicePage() {
	browser := rod.New().MustConnect().MustPage("")
	downloadPath := os.Args[1]
	proto.PageSetDownloadBehavior{
		Behavior:     proto.PageSetDownloadBehaviorBehaviorAllow,
		DownloadPath: downloadPath,
	}.Call(browser)
	c.Page = browser.Browser().MustPage(invoiceUrl)
}

func (c *crawler) mustFillInvoiceData() {
	c.MustElement(`[name="invoiceCode"]`).MustInput(strconv.Itoa(c.invoiceData.InvoiceCode))
	c.MustElement("#companieName").MustInput(c.invoiceData.Provider)
	c.mustTypeSlowly(`[data-testid="cnpj"]`, c.invoiceData.CNPJ)
	c.MustElement("#address").MustInput(c.invoiceData.Address)
	c.MustElement("#client-email").MustInput(c.invoiceData.Email)
	c.MustElement("#customerName").MustInput(c.invoiceData.Client)
	c.MustElement(`[name="customerCnpjCpf"]`).MustInput(c.invoiceData.ClientCNPJ)
	c.MustElement(`button[for="opening-balance"]`).MustClick()
	c.selectCurrency()
	c.MustElement("#opening-balance").MustInput(c.invoiceData.Amount)
	c.MustElement("#fatura-vencimento").MustInput(c.invoiceData.DueDate)
	c.MustElement(`[name="serviceTitle"]`).MustInput(c.invoiceData.ServiceTitle)
	c.MustElement("#description").MustInput(c.invoiceData.ServiceDescription)
	c.MustElement(`[name="swiftCode"]`).MustInput(c.invoiceData.Swift)
	c.MustElement(`[name="ibanCode"]`).MustInput(c.invoiceData.Iban)
}

func (c *crawler) mustTypeSlowly(selector, text string) {
	element := c.MustElement(selector)
	for _, key := range text {
		time.Sleep(time.Millisecond * 100)
		element.MustInput(string(key))
	}
	time.Sleep(time.Millisecond * 100)
}

func (c *crawler) selectCurrency() {
	currencyList := c.MustElement("#prefix-dropdown")
	currencies := currencyList.MustElements("li")
	for _, currency := range currencies {
		if currency.MustText() == c.invoiceData.Currency {
			currency.MustClick()
			break
		}
	}

}

func (c *crawler) clickDownload() (*rod.Element, error) {
	pageButtons := c.MustElements(`button`)
	for _, pageButton := range pageButtons {
		if pageButton.MustText() == "Baixar invoice" {
			return pageButton.MustClick(), nil
		}
	}
	return nil, errors.New("download button not found")
}

func (c *crawler) Run() error {
	fmt.Println("Filling invoice form")
	c.mustFillInvoiceData()
	fmt.Println("Downloading invoice")
	_, err := c.clickDownload()
	if err != nil {
		return fmt.Errorf("crawler.Run: %w", err)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Closing")
	return c.Close()
}

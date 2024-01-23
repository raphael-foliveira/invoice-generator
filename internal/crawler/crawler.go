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

var invoiceUrl = "https://invoice.agilize.com.br/"

type crawler struct {
	invoiceData  InvoiceData
	downloadPath string
	page         *rod.Page
}

func NewCrawler(invoiceData *InvoiceData, downloadPath string) *crawler {
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
	c.page = browser.Browser().MustPage(invoiceUrl)
}

func (c *crawler) mustFillInvoiceData() {
	c.page.MustElement(`[name="invoiceCode"]`).MustInput(strconv.Itoa(c.invoiceData.InvoiceCode))
	c.page.MustElement("#companieName").MustInput(c.invoiceData.Provider)
	c.page.MustElement(`[data-testid="cnpj"]`).MustInput(c.invoiceData.CNPJ)
	c.page.MustElement("#address").MustInput(c.invoiceData.Address)
	c.page.MustElement("#client-email").MustInput(c.invoiceData.Email)
	c.page.MustElement("#customerName").MustInput(c.invoiceData.Client)
	c.page.MustElement(`[name="customerCnpjCpf"]`).MustInput(c.invoiceData.ClientCNPJ)
	c.page.MustElement(`button[for="opening-balance"]`).MustClick()
	currencyList := c.page.MustElement("#prefix-dropdown")
	currencies := currencyList.MustElements("li")
	for _, currency := range currencies {
		if currency.MustText() == c.invoiceData.Currency {
			currency.MustClick()
			break
		}
	}
	c.page.MustElement("#opening-balance").MustInput(c.invoiceData.Amount)
	c.page.MustElement("#fatura-vencimento").MustInput(c.invoiceData.DueDate)
	c.page.MustElement(`[name="serviceTitle"]`).MustInput(c.invoiceData.ServiceTitle)
	c.page.MustElement("#description").MustInput(c.invoiceData.ServiceDescription)
	c.page.MustElement(`[name="swiftCode"]`).MustInput(c.invoiceData.Swift)
	c.page.MustElement(`[name="ibanCode"]`).MustInput(c.invoiceData.Iban)
}

func (c *crawler) clickDownload() (*rod.Element, error) {
	pageButtons := c.page.MustElements(`button`)
	for _, pageButton := range pageButtons {
		if pageButton.MustText() == "Baixar invoice" {
			return pageButton.MustClick(), nil
		}
	}
	return nil, errors.New("Download button not found")
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
	return c.page.Close()
}

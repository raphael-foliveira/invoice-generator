package crawler

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
	Swift              string `json:"swift"`
	Iban               string `json:"iban"`
}

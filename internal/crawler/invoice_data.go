package crawler

type InvoiceData struct {
	ClientCNPJ         string `json:"clientCnpj"`
	Amount             string `json:"amount"`
	CNPJ               string `json:"cnpj"`
	Address            string `json:"address"`
	Email              string `json:"email"`
	Client             string `json:"client"`
	Provider           string `json:"provider"`
	Iban               string `json:"iban"`
	ServiceTitle       string `json:"serviceTitle"`
	DueDate            string `json:"dueDate"`
	Currency           string `json:"currency"`
	ServiceDescription string `json:"serviceDescription"`
	Swift              string `json:"swift"`
	InvoiceCode        int    `json:"invoiceCode"`
}

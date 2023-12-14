package aycd

type Profile struct {
	Name                          string         `json:"name"`
	Size                          string         `json:"size"`
	ProfileGroup                  string         `json:"profileGroup"`
	Notes                         string         `json:"notes"`
	BillingAddress                Address        `json:"billingAddress"`
	ShippingAddress               Address        `json:"shippingAddress"`
	PaymentDetails                PaymentDetails `json:"paymentDetails"`
	SameBillingAndShippingAddress bool           `json:"sameBillingAndShippingAddress"`
	OnlyCheckoutOnce              bool           `json:"onlyCheckoutOnce"`
	MatchNameOnCardAndAddress     bool           `json:"matchNameOnCardAndAddress"`
}

type Address struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Line3    string `json:"line3"`
	PostCode string `json:"postCode"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
}

type PaymentDetails struct {
	NameOnCard   string `json:"nameOnCard"`
	CardType     string `json:"cardType"`
	CardNumber   string `json:"cardNumber"`
	CardExpMonth string `json:"cardExpMonth"`
	CardExpYear  string `json:"cardExpYear"`
	CardCVV      string `json:"cardCvv"`
}

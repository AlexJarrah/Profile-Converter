package stellar

type Profile struct {
	ProfileName       string   `json:"profileName"`
	Email             string   `json:"email"`
	Phone             string   `json:"phone"`
	Shipping          Shipping `json:"shipping"`
	BillingAsShipping bool     `json:"billingAsShipping"`
	Billing           Billing  `json:"billing"`
	Payment           Payment  `json:"payment"`
}

type Shipping struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Address   string `json:"address"`
	Address2  string `json:"address2"`
	State     string `json:"state"`
	City      string `json:"city"`
	Zipcode   string `json:"zipcode"`
}

type Billing struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Address   string `json:"address"`
	Address2  string `json:"address2"`
	State     string `json:"state"`
	City      string `json:"city"`
	Zipcode   string `json:"zipcode"`
}

type Payment struct {
	CardName   string `json:"cardName"`
	CardType   string `json:"cardType"`
	CardNumber string `json:"cardNumber"`
	CardMonth  string `json:"cardMonth"`
	CardYear   string `json:"cardYear"`
	CardCVV    string `json:"cardCvv"`
}

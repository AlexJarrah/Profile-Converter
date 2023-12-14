package internal

type ProfileGroup struct {
	Name     string
	Profiles []Profile
}
type Profile struct {
	ProfileName       string  `json:"profileName"`
	Email             string  `json:"email"`
	Phone             string  `json:"phone"`
	Shipping          Address `json:"shipping"`
	BillingAsShipping bool    `json:"billingAsShipping"`
	Billing           Address `json:"billing"`
	Payment           Payment `json:"payment"`
}
type Address struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Address   string `json:"address"`
	Address2  string `json:"address2"`
	State     string `json:"state"`
	City      string `json:"city"`
	Zipcode   uint32 `json:"zipcode"`
}
type Payment struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Number uint64 `json:"number"`
	Month  uint8  `json:"month"`
	Year   uint16 `json:"year"`
	CVV    uint16 `json:"cvv"`
}

type Format int

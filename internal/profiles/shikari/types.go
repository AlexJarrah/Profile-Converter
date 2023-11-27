package shikari

type Profiles struct {
	Profiles []Profile
}

type Profile struct {
	ProfileName      string `csv:"profile_name"`
	FirstName        string `csv:"first_name"`
	LastName         string `csv:"last_name"`
	Email            string `csv:"email"`
	PhoneNum         string `csv:"phone_num"`
	CCNumber         string `csv:"cc_number"`
	CCExpMonth       string `csv:"cc_exp_month"`
	CCExpYear        string `csv:"cc_exp_year"`
	CCCVV            string `csv:"cc_cvv"`
	ShippingStreet   string `csv:"shipping_street"`
	ShippingStreet2  string `csv:"shipping_street_2"`
	ShippingCity     string `csv:"shipping_city"`
	ShippingState    string `csv:"shipping_state"`
	ShippingZipCode  string `csv:"shipping_zip_code"`
	ShippingCountry  string `csv:"shipping_country"`
	BillingFirstName string `csv:"billing_first_name"`
	BillingLastName  string `csv:"billing_last_name"`
	BillingStreet    string `csv:"billing_street"`
	BillingStreet2   string `csv:"billing_street_2"`
	BillingCity      string `csv:"billing_city"`
	BillingState     string `csv:"billing_state"`
	BillingZipCode   string `csv:"billing_zip_code"`
	BillingCountry   string `csv:"billing_country"`
}

package shikari

import (
	"fmt"

	"github.com/AlexJarrah/Profile-Converter/internal"
)

// CSV header
const header = "profile_name,first_name,last_name,email,phone_num,cc_number,cc_exp_month,cc_exp_year,cc_cvv,shipping_street,shipping_street_2,shipping_city,shipping_state,shipping_zip_code,shipping_country,billing_first_name,billing_last_name,billing_street,billing_street_2,billing_city,billing_state,billing_zip_code,billing_country"

// Converts the universal struct into the Stellar struct.
// Returns the result as a struct and as CSV.
func Convert(profiles []internal.Profile) (res []Profile, resCSV string, err error) {
	for _, p := range profiles {
		profile := Profile{
			ProfileName:     p.ProfileName,
			FirstName:       p.Shipping.FirstName,
			LastName:        p.Shipping.LastName,
			Email:           p.Email,
			PhoneNum:        p.Phone,
			CCNumber:        fmt.Sprint(p.Payment.Number),
			CCExpMonth:      fmt.Sprint(p.Payment.Month),
			CCExpYear:       fmt.Sprint(p.Payment.Year),
			CCCVV:           fmt.Sprint(p.Payment.CVV),
			ShippingStreet:  p.Shipping.Address,
			ShippingStreet2: p.Shipping.Address2,
			ShippingCity:    p.Shipping.City,
			ShippingState:   p.Shipping.State,
			ShippingZipCode: fmt.Sprint(p.Shipping.Zipcode),
			ShippingCountry: p.Shipping.Country,
		}

		if !p.BillingAsShipping {
			profile.BillingFirstName = p.Billing.FirstName
			profile.BillingLastName = p.Billing.LastName
			profile.BillingStreet = p.Billing.Address
			profile.BillingStreet2 = p.Billing.Address2
			profile.BillingCity = p.Billing.City
			profile.BillingState = p.Billing.State
			profile.BillingZipCode = fmt.Sprint(p.Billing.Zipcode)
			profile.BillingCountry = p.Billing.Country
		}

		res = append(res, profile)
	}

	// Creates the CSV value
	resCSV = header + "\n"
	for _, p := range res {
		resCSV += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			p.ProfileName,
			p.FirstName,
			p.LastName,
			p.Email,
			p.PhoneNum,
			p.CCNumber,
			p.CCExpMonth,
			p.CCExpYear,
			p.CCCVV,
			p.ShippingStreet,
			p.ShippingStreet2,
			p.ShippingCity,
			p.ShippingState,
			p.ShippingZipCode,
			p.ShippingCountry,
			p.BillingFirstName,
			p.BillingLastName,
			p.BillingStreet,
			p.BillingStreet2,
			p.BillingCity,
			p.BillingState,
			p.BillingZipCode,
			p.BillingCountry)
	}

	return res, resCSV, err
}

package aycd

import (
	"encoding/json"
	"fmt"

	"github.com/AlexJarrah/Profile-Converter/internal"
)

// Converts the universal struct into the AYCD struct.
// Returns the result as a struct and as JSON.
func Convert(profiles []internal.Profile) (res []Profile, resJSON string, err error) {
	for _, p := range profiles {
		profile := Profile{
			Name:         p.ProfileName,
			Size:         "",
			ProfileGroup: "",
			Notes:        "",
			BillingAddress: Address{
				Name:     fmt.Sprintf("%s %s", p.Billing.FirstName, p.Billing.LastName),
				Email:    p.Email,
				Phone:    p.Phone,
				Line1:    p.Billing.Address,
				Line2:    p.Billing.Address2,
				Line3:    "",
				PostCode: fmt.Sprint(p.Billing.Zipcode),
				City:     p.Billing.City,
				Country:  internal.FormatCountry(p.Billing.Country, true),
				State:    internal.FormatState(p.Billing.State, true),
			},
			ShippingAddress: Address{
				Name:     fmt.Sprintf("%s %s", p.Shipping.FirstName, p.Shipping.LastName),
				Email:    p.Email,
				Phone:    p.Phone,
				Line1:    p.Shipping.Address,
				Line2:    p.Shipping.Address2,
				Line3:    "",
				PostCode: fmt.Sprint(p.Shipping.Zipcode),
				City:     p.Shipping.City,
				Country:  internal.FormatCountry(p.Shipping.Country, true),
				State:    internal.FormatState(p.Shipping.State, true),
			},
			PaymentDetails: PaymentDetails{
				NameOnCard:   fmt.Sprintf("%s %s", p.Shipping.FirstName, p.Shipping.LastName),
				CardType:     p.Payment.Type,
				CardNumber:   fmt.Sprint(p.Payment.Number),
				CardExpMonth: internal.FormatCardMonth(p.Payment.Month, true),
				CardExpYear:  fmt.Sprint(internal.FormatCardYear(p.Payment.Year, true)),
				CardCVV:      fmt.Sprint(p.Payment.CVV),
			},
			SameBillingAndShippingAddress: p.BillingAsShipping,
			OnlyCheckoutOnce:              false,
			MatchNameOnCardAndAddress:     false,
		}

		if p.Payment.Name == fmt.Sprintf("%s %s", p.Shipping.FirstName, p.Shipping.LastName) {
			profile.MatchNameOnCardAndAddress = true
		}

		res = append(res, profile)
	}

	resJSONBytes, err := json.Marshal(res)
	return res, string(resJSONBytes), err
}

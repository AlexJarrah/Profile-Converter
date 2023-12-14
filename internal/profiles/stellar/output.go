package stellar

import (
	"encoding/json"
	"fmt"

	"github.com/AlexJarrah/Profile-Converter/internal"
)

// Converts the universal struct into the Stellar struct.
// Returns the result as a struct and as JSON.
func Convert(profiles []internal.Profile) (res []Profile, resJSON string, err error) {
	for _, p := range profiles {
		profile := Profile{
			ProfileName: p.ProfileName,
			Email:       p.Email,
			Phone:       p.Phone,
			Shipping: Shipping{
				FirstName: p.Shipping.FirstName,
				LastName:  p.Shipping.LastName,
				Country:   internal.FormatCountry(p.Shipping.Country, false),
				Address:   p.Shipping.Address,
				Address2:  p.Shipping.Address2,
				State:     internal.FormatState(p.Shipping.State, false),
				City:      p.Shipping.City,
				Zipcode:   fmt.Sprint(p.Shipping.Zipcode),
			},
			BillingAsShipping: p.BillingAsShipping,
			Billing: Billing{
				FirstName: p.Shipping.FirstName,
				LastName:  p.Shipping.LastName,
				Country:   internal.FormatCountry(p.Shipping.Country, false),
				Address:   p.Shipping.Address,
				Address2:  p.Shipping.Address2,
				State:     internal.FormatState(p.Shipping.State, false),
				City:      p.Shipping.City,
				Zipcode:   fmt.Sprint(p.Shipping.Zipcode),
			},
			Payment: Payment{
				CardName:   p.Payment.Name,
				CardType:   p.Payment.Type,
				CardNumber: fmt.Sprint(p.Payment.Number),
				CardMonth:  fmt.Sprint(internal.FormatCardMonth(p.Payment.Month, true)),
				CardYear:   fmt.Sprint(internal.FormatCardYear(p.Payment.Year, false)),
				CardCVV:    fmt.Sprint(p.Payment.CVV),
			},
		}

		res = append(res, profile)
	}

	resJSONBytes, err := json.Marshal(res)
	return res, string(resJSONBytes), err
}

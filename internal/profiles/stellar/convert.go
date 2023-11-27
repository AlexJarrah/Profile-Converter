package stellar

import (
	"fmt"

	"github.com/quo0001/Profile-Converter/internal"
)

func Convert(profiles []internal.Profile) []Profile {
	resp := []Profile{}

	for _, p := range profiles {
		resp = append(resp, Profile{
			ProfileName: p.ProfileName,
			Email:       p.Email,
			Phone:       p.Phone,
			Shipping: Shipping{
				FirstName: p.Shipping.FirstName,
				LastName:  p.Shipping.LastName,
				Country:   p.Shipping.Country,
				Address:   p.Shipping.Address,
				Address2:  p.Shipping.Address2,
				State:     p.Shipping.State,
				City:      p.Shipping.City,
				Zipcode:   fmt.Sprint(p.Shipping.Zipcode),
			},
			BillingAsShipping: p.BillingAsShipping,
			Billing: Billing{
				FirstName: p.Shipping.FirstName,
				LastName:  p.Shipping.LastName,
				Country:   p.Shipping.Country,
				Address:   p.Shipping.Address,
				Address2:  p.Shipping.Address2,
				State:     p.Shipping.State,
				City:      p.Shipping.City,
				Zipcode:   fmt.Sprint(p.Shipping.Zipcode),
			},
			Payment: Payment{
				CardName:   p.Payment.Name,
				CardType:   p.Payment.Type,
				CardNumber: fmt.Sprint(p.Payment.Number),
				CardMonth:  fmt.Sprint(p.Payment.Month),
				CardYear:   fmt.Sprint(p.Payment.Year),
				CardCvv:    fmt.Sprint(p.Payment.CVV),
			},
		})
	}

	return resp
}

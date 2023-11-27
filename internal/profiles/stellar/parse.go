package stellar

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/quo0001/Profile-Converter/internal"
)

func Parse(path string) ([]internal.Profile, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	if err = json.Unmarshal(f, &profiles); err != nil {
		return nil, err
	}

	resp := []internal.Profile{}

	for _, p := range profiles {
		zip, _ := strconv.Atoi(p.Shipping.Zipcode)
		num, _ := strconv.Atoi(p.Payment.CardNumber)
		month, _ := strconv.Atoi(p.Payment.CardMonth)
		year, _ := strconv.Atoi(p.Payment.CardYear)
		year += 2000
		cvv, _ := strconv.Atoi(p.Payment.CardCvv)

		profile := internal.Profile{
			ProfileName: p.ProfileName,
			Email:       p.Email,
			Phone:       p.Phone,
			Shipping: internal.Address{
				FirstName: p.Shipping.FirstName,
				LastName:  p.Shipping.LastName,
				Country:   p.Shipping.Country,
				Address:   p.Shipping.Address,
				Address2:  p.Shipping.Address2,
				State:     p.Shipping.State,
				City:      p.Shipping.City,
				Zipcode:   uint32(zip),
			},
			BillingAsShipping: p.BillingAsShipping,
			Billing: internal.Address{
				FirstName: p.Billing.FirstName,
				LastName:  p.Billing.LastName,
				Country:   p.Billing.Country,
				Address:   p.Billing.Address,
				Address2:  p.Billing.Address2,
				State:     p.Billing.State,
				City:      p.Billing.City,
				Zipcode:   uint32(zip),
			},
			Payment: internal.Payment{
				Name:   p.Payment.CardName,
				Type:   p.Payment.CardType,
				Number: uint64(num),
				Month:  uint8(month),
				Year:   uint16(year),
				CVV:    uint16(cvv),
			},
		}

		resp = append(resp, profile)
	}

	return resp, nil
}

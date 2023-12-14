package stellar

import (
	"encoding/json"
	"os"

	"github.com/AlexJarrah/Profile-Converter/internal"
)

// Parses a file's contents into the universal struct format
func Parse(path string) ([]internal.Profile, error) {
	// Reads the specified file path
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshals the JSON file
	var profiles []Profile
	if err = json.Unmarshal(f, &profiles); err != nil {
		return nil, err
	}

	// Initialize the result variable
	var result []internal.Profile

	// Iterates through each profile and appends it to the result
	for _, p := range profiles {
		profile := internal.Profile{
			ProfileName: p.ProfileName,
			Email:       p.Email,
			Phone:       p.Phone,
			Shipping: internal.Address{
				FirstName: p.Shipping.FirstName,
				LastName:  p.Shipping.LastName,
				Country:   internal.FormatCountry(p.Shipping.Country, true),
				Address:   p.Shipping.Address,
				Address2:  p.Shipping.Address2,
				State:     internal.FormatState(p.Shipping.State, true),
				City:      p.Shipping.City,
				Zipcode:   uint32(internal.ParseInt(p.Shipping.Zipcode)),
			},
			BillingAsShipping: p.BillingAsShipping,
			Billing: internal.Address{
				FirstName: p.Billing.FirstName,
				LastName:  p.Billing.LastName,
				Country:   internal.FormatCountry(p.Billing.Country, true),
				Address:   p.Billing.Address,
				Address2:  p.Billing.Address2,
				State:     internal.FormatState(p.Billing.State, true),
				City:      p.Billing.City,
				Zipcode:   uint32(internal.ParseInt(p.Shipping.Zipcode)),
			},
			Payment: internal.Payment{
				Name:   p.Payment.CardName,
				Type:   p.Payment.CardType,
				Number: uint64(internal.ParseInt(p.Payment.CardNumber)),
				Month:  uint8(internal.ParseInt(internal.FormatCardMonth(p.Payment.CardMonth, true))),
				Year:   internal.FormatCardYear(p.Payment.CardYear, true),
				CVV:    uint16(internal.ParseInt(p.Payment.CardCVV)),
			},
		}

		// Appends the profile to the result
		result = append(result, profile)
	}

	return result, nil
}

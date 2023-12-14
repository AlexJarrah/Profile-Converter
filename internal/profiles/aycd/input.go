package aycd

import (
	"encoding/json"
	"os"
	"strings"

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

		// Splits the name into first and last
		name := strings.Split(p.ShippingAddress.Name, " ")
		sFirstName := name[0]
		var sLastName string
		if len(name) > 1 {
			sLastName = name[1]
		}

		name = strings.Split(p.BillingAddress.Name, " ")
		bFirstName := name[0]
		var bLastName string
		if len(name) > 1 {
			bLastName = name[1]
		}

		profile := internal.Profile{
			ProfileName: p.Name,
			Email:       p.BillingAddress.Email,
			Phone:       p.BillingAddress.Phone,
			Shipping: internal.Address{
				FirstName: sFirstName,
				LastName:  sLastName,
				Country:   internal.FormatCountry(p.ShippingAddress.Country, true),
				Address:   p.ShippingAddress.Line1,
				Address2:  p.ShippingAddress.Line2,
				State:     internal.FormatState(p.ShippingAddress.State, true),
				City:      p.ShippingAddress.City,
				Zipcode:   uint32(internal.ParseInt(p.ShippingAddress.PostCode)),
			},
			BillingAsShipping: p.SameBillingAndShippingAddress,
			Billing: internal.Address{
				FirstName: bFirstName,
				LastName:  bLastName,
				Country:   internal.FormatCountry(p.BillingAddress.Country, true),
				Address:   p.BillingAddress.Line1,
				Address2:  p.BillingAddress.Line2,
				State:     internal.FormatState(p.BillingAddress.State, true),
				City:      p.BillingAddress.City,
				Zipcode:   uint32(internal.ParseInt(p.BillingAddress.PostCode)),
			},
			Payment: internal.Payment{
				Name:   p.PaymentDetails.NameOnCard,
				Type:   p.PaymentDetails.CardType,
				Number: uint64(internal.ParseInt(p.PaymentDetails.CardNumber)),
				Month:  uint8(internal.ParseInt(internal.FormatCardMonth(p.PaymentDetails.CardExpMonth, true))),
				Year:   internal.FormatCardYear(p.PaymentDetails.CardExpYear, true),
				CVV:    uint16(internal.ParseInt(p.PaymentDetails.CardCVV)),
			},
		}

		// Appends the profile to the result
		result = append(result, profile)
	}

	return result, nil
}

package shikari

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/AlexJarrah/Profile-Converter/internal"
)

// Parses a file's contents into the universal struct format
func Parse(path string) ([]internal.Profile, error) {
	// Reads the specified file path
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Unmarshals the CSV file
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Initialize the result variable
	var result []internal.Profile

	// Iterates through each profile and appends it to the result
	for i, r := range records {
		// Ignore the first record (header)
		if i == 0 {
			continue
		}

		profile := internal.Profile{
			ProfileName: r[0],
			Email:       r[3],
			Phone:       r[4],
			Shipping: internal.Address{
				FirstName: r[1],
				LastName:  r[2],
				Country:   internal.FormatCountry(r[14], true),
				Address:   r[9],
				Address2:  r[10],
				State:     internal.FormatState(r[12], true),
				City:      r[11],
				Zipcode:   uint32(internal.ParseInt(r[13])),
			},
			BillingAsShipping: false,
			Billing: internal.Address{
				FirstName: r[15],
				LastName:  r[16],
				Country:   internal.FormatCountry(r[22], true),
				Address:   r[17],
				Address2:  r[18],
				State:     internal.FormatState(r[20], true),
				City:      r[19],
				Zipcode:   uint32(internal.ParseInt(r[21])),
			},
			Payment: internal.Payment{
				Name:   fmt.Sprintf("%s %s", r[1], r[2]),
				Type:   internal.GetCardType(r[5]),
				Number: uint64(internal.ParseInt(r[5])),
				Month:  uint8(internal.ParseInt(internal.FormatCardMonth(r[6], true))),
				Year:   internal.FormatCardYear(r[7], true),
				CVV:    uint16(internal.ParseInt(r[8])),
			},
		}

		// Check if the billing address is undefined or the same as the shipping
		if profile.Billing == (internal.Address{}) || profile.Billing == profile.Shipping {
			profile.BillingAsShipping = true
		}

		// Appends the profile to the result
		result = append(result, profile)
	}

	return result, nil
}

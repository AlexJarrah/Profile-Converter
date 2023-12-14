package trickle

import (
	"encoding/csv"
	"os"
	"strings"

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

		// Splits the name into first and last
		name := strings.Split(r[1], " ")
		firstName := name[0]
		var lastName string
		if len(name) > 1 {
			lastName = name[1]
		}

		profile := internal.Profile{
			ProfileName: "",
			Email:       r[4],
			Phone:       r[5],
			Shipping: internal.Address{
				FirstName: firstName,
				LastName:  lastName,
				Country:   internal.FormatCountry(r[11], true),
				Address:   r[6],
				Address2:  r[7],
				State:     internal.FormatState(r[8], true),
				City:      r[9],
				Zipcode:   uint32(internal.ParseInt(r[10])),
			},
			BillingAsShipping: true,
			Billing: internal.Address{
				FirstName: firstName,
				LastName:  lastName,
				Country:   internal.FormatCountry(r[11], true),
				Address:   r[6],
				Address2:  r[7],
				State:     internal.FormatState(r[8], true),
				City:      r[9],
				Zipcode:   uint32(internal.ParseInt(r[10])),
			},
			Payment: internal.Payment{
				Name:   r[2],
				Type:   internal.GetCardType(r[12]),
				Number: uint64(internal.ParseInt(r[12])),
				Month:  uint8(internal.ParseInt(internal.FormatCardMonth(r[13], true))),
				Year:   internal.FormatCardYear(r[14], true),
				CVV:    uint16(internal.ParseInt(r[15])),
			},
		}

		// Appends the profile to the result
		result = append(result, profile)
	}

	return result, nil
}

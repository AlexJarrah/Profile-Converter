package trickle

import (
	"encoding/csv"
	"os"
	"strconv"
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

		// Parses strings into integers
		zip, _ := strconv.Atoi(r[10])
		num, _ := strconv.Atoi(r[12])
		month, _ := strconv.Atoi(r[13])
		year, _ := strconv.Atoi(r[14])
		cvv, _ := strconv.Atoi(r[15])

		profile := internal.Profile{
			ProfileName: "",
			Email:       r[4],
			Phone:       r[5],
			Shipping: internal.Address{
				FirstName: firstName,
				LastName:  lastName,
				Country:   r[11],
				Address:   r[6],
				Address2:  r[7],
				State:     r[8],
				City:      r[10],
				Zipcode:   uint32(zip),
			},
			BillingAsShipping: true,
			Billing: internal.Address{
				FirstName: firstName,
				LastName:  lastName,
				Country:   r[11],
				Address:   r[6],
				Address2:  r[7],
				State:     r[8],
				City:      r[10],
				Zipcode:   uint32(zip),
			},
			Payment: internal.Payment{
				Name:   r[2],
				Type:   internal.GetCardType(uint64(num)),
				Number: uint64(num),
				Month:  uint8(month),
				Year:   uint16(year),
				CVV:    uint16(cvv),
			},
		}

		// Appends the profile to the result
		result = append(result, profile)
	}

	return result, nil
}

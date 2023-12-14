package shikari

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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

		// Parses strings into integers
		shippingZip, _ := strconv.Atoi(r[13])
		billingZip, _ := strconv.Atoi(r[21])
		num, _ := strconv.Atoi(r[5])
		month, _ := strconv.Atoi(r[6])
		year, _ := strconv.Atoi(r[7])
		cvv, _ := strconv.Atoi(r[8])

		profile := internal.Profile{
			ProfileName: r[0],
			Email:       r[3],
			Phone:       r[4],
			Shipping: internal.Address{
				FirstName: r[1],
				LastName:  r[2],
				Country:   r[14],
				Address:   r[9],
				Address2:  r[10],
				State:     r[12],
				City:      r[11],
				Zipcode:   uint32(shippingZip),
			},
			BillingAsShipping: false,
			Billing: internal.Address{
				FirstName: r[15],
				LastName:  r[16],
				Country:   r[22],
				Address:   r[17],
				Address2:  r[18],
				State:     r[20],
				City:      r[19],
				Zipcode:   uint32(billingZip),
			},
			Payment: internal.Payment{
				Name:   fmt.Sprintf("%s %s", r[1], r[2]),
				Type:   internal.GetCardType(uint64(num)),
				Number: uint64(num),
				Month:  uint8(month),
				Year:   uint16(year),
				CVV:    uint16(cvv),
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

package shikari

import (
	"encoding/csv"
	"os"
)

func Parse(path string) (Profiles, error) {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return Profiles{}, err
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Profiles{}, err
	}

	// Create a Profiles instance to hold the parsed data
	var profiles Profiles

	// Iterate through the records and convert them to Profile structs
	for i, record := range records {
		if i == 0 {
			continue
		}
		profile := Profile{
			ProfileName:      record[0],
			FirstName:        record[1],
			LastName:         record[2],
			Email:            record[3],
			PhoneNum:         record[4],
			CCNumber:         record[5],
			CCExpMonth:       record[6],
			CCExpYear:        record[7],
			CCCVV:            record[8],
			ShippingStreet:   record[9],
			ShippingStreet2:  record[10],
			ShippingCity:     record[11],
			ShippingState:    record[12],
			ShippingZipCode:  record[13],
			ShippingCountry:  record[14],
			BillingFirstName: record[15],
			BillingLastName:  record[16],
			BillingStreet:    record[17],
			BillingStreet2:   record[18],
			BillingCity:      record[19],
			BillingState:     record[20],
			BillingZipCode:   record[21],
			BillingCountry:   record[22],
		}

		// Append the Profile instance to the Profiles slice
		profiles.Profiles = append(profiles.Profiles, profile)
	}

	return profiles, nil
}

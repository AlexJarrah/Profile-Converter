package shikari

import (
	"strconv"

	"github.com/quo0001/Profile-Converter/internal"
)

func (profiles Profiles) Convert() []internal.Profile {
	var result []internal.Profile

	for _, p := range profiles.Profiles {
		shipzip, _ := strconv.Atoi(p.ShippingZipCode)
		bilsip, _ := strconv.Atoi(p.BillingZipCode)
		num, month, year, cvv := atoi(p.CCNumber, p.CCExpMonth, p.CCExpYear, p.CCCVV)

		billing := internal.Address{
			FirstName: p.BillingFirstName,
			LastName:  p.BillingLastName,
			Country:   p.BillingCountry,
			Address:   p.BillingStreet,
			Address2:  p.BillingStreet2,
			State:     p.BillingState,
			City:      p.BillingCity,
			Zipcode:   uint32(bilsip),
		}

		profile := internal.Profile{
			ProfileName: p.ProfileName,
			Email:       p.Email,
			Phone:       p.PhoneNum,
			Shipping: internal.Address{
				FirstName: p.FirstName,
				LastName:  p.LastName,
				Country:   p.ShippingCountry,
				Address:   p.ShippingStreet,
				Address2:  p.ShippingStreet2,
				State:     p.ShippingState,
				City:      p.ShippingCity,
				Zipcode:   uint32(shipzip),
			},
			BillingAsShipping: p.BillingStreet == "",
			Billing:           billing,
			Payment: internal.Payment{
				Name:   p.CCNumber,
				Number: uint64(num),
				Month:  uint8(month),
				Year:   uint16(year),
				CVV:    uint16(cvv),
			},
		}

		if profile.BillingAsShipping {
			profile.Billing = profile.Shipping
		}

		result = append(result, profile)
	}

	return result
}

func atoi(strs ...string) (int, int, int, int) {
	var result [4]int

	for i, str := range strs {
		result[i], _ = strconv.Atoi(str)
	}

	return result[0], result[1], result[2], result[3]
}

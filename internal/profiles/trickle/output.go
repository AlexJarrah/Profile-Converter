package trickle

import (
	"fmt"

	"github.com/AlexJarrah/Profile-Converter/internal"
)

// CSV header
const header = "Keyword,SIZE/QUANTITY,FIRST NAME,LAST NAME,EMAIL,PHONE NUMBER,ADDRESS 1,ADDRESS 2,STATE,CITY,ZIP,COUNTRY,CC NUMBER,MONTH,YEAR,CVC,Task Quantity,Retry Delay,Monitor Delay,Site,Mode,2CaptchaKEY,SHIPPINGRATE"

// Converts the universal struct into the Trickle struct.
// Returns the result as a struct and as CSV.
func Convert(profiles []internal.Profile) (res []Profile, resCSV string, err error) {
	for _, p := range profiles {
		profile := Profile{
			Keyword:      `""`,
			Size:         "",
			FirstName:    p.Billing.FirstName,
			LastName:     p.Billing.LastName,
			Email:        p.Email,
			PhoneNumber:  p.Phone,
			Address1:     p.Billing.Address,
			Address2:     p.Billing.Address2,
			State:        internal.FormatState(p.Billing.State, false),
			City:         p.Billing.City,
			Zip:          fmt.Sprint(p.Billing.Zipcode),
			Country:      internal.FormatCountry(p.Billing.Country, false),
			CCNumber:     fmt.Sprint(p.Payment.Number),
			Month:        internal.FormatCardMonth(fmt.Sprint(p.Payment.Month), true),
			Year:         fmt.Sprint(internal.FormatCardYear(p.Payment.Year, true)),
			CVC:          fmt.Sprint(p.Payment.CVV),
			TaskQuantity: "",
			RetryDelay:   "",
			MonitorDelay: "",
			Site:         "",
			Mode:         "",
			_2CaptchaKEY: "",
			SHIPPINGRATE: "",
		}

		res = append(res, profile)
	}

	// Creates the CSV value
	resCSV = header + "\n"
	for _, p := range res {
		resCSV += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", p.Keyword, p.Size, p.FirstName, p.LastName, p.Email, p.PhoneNumber, p.Address1, p.Address2, p.State, p.City, p.Zip, p.Country, p.CCNumber, p.Month, p.Year, p.CVC, p.TaskQuantity, p.RetryDelay, p.MonitorDelay, p.Site, p.Mode, p._2CaptchaKEY, p.SHIPPINGRATE)
	}

	return res, resCSV, err
}

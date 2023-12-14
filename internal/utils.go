package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCardType(CardNumber any) string {
	number := fmt.Sprint(CardNumber)

	if len(number) != 16 {
		return ""
	}

	digit1 := string(number[0])
	switch digit1 {
	case "3":
		digit2 := string(number[1])
		if digit2 == "4" || digit2 == "7" {
			return "American Express"
		}
	case "4":
		return "Visa"
	case "5":
		return "Mastercard"
	case "6":
		return "Discover"
	}

	return ""
}

func FormatCardYear(CardYear any, full bool) uint16 {
	year := fmt.Sprint(CardYear)

	switch full {
	case true:
		if len(year) == 2 {
			year = "20" + year
		}

	case false:
		if len(year) == 4 {
			year = year[2:]
		}
	}

	res, _ := strconv.Atoi(year)
	return uint16(res)
}

func FormatCardMonth(CardMonth any, full bool) string {
	month := fmt.Sprint(CardMonth)

	switch full {
	case true:
		if len(month) == 1 {
			month = "0" + month
		}

	case false:
		month = strings.TrimLeft(month, "0")
	}

	return month
}

func FormatCountry(country string, full bool) string {
	country = strings.ToUpper(country)

	switch full {
	case true:
		if long, exists := CountryNames[country]; exists {
			country = long
		}

	case false:
		for short, long := range CountryNames {
			if strings.ToUpper(long) == country {
				country = short
			}
		}
	}

	return country
}

func FormatState(state string, full bool) string {
	state = strings.ToUpper(state)

	switch full {
	case true:
		if long, exists := StateNames[state]; exists {
			state = long
		}

	case false:
		for short, long := range StateNames {
			if strings.ToUpper(long) == state {
				state = short
			}
		}
	}

	return state
}

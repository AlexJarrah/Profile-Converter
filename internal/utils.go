package internal

import "fmt"

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

package main

import (
	"fmt"
	"log"

	"github.com/quo0001/Profile-Converter/internal"
	"github.com/quo0001/Profile-Converter/internal/profiles/shikari"
	"github.com/quo0001/Profile-Converter/internal/profiles/stellar"
)

func main() {
	// path, inputFormat, outputFormat := prompt.Prompt()
	path := "formats/profiles/stellar.json"
	inputFormat := internal.FormatStellar
	outputFormat := internal.FormatShikari

	var profiles []internal.Profile
	var err error

	switch inputFormat {
	case internal.FormatStellar:
		profiles, err = stellar.Parse(path)

	case internal.FormatShikari:
		profiles, err = shikari.Parse(path)
	}

	if err != nil {
		log.Println("error:", err)
		return
	}

	fmt.Println(profiles[0].Billing.Address)

	switch outputFormat {
	case internal.FormatStellar:
		profiles, output, err := stellar.Convert(profiles)
		if err != nil {
			log.Println("error:", err)
			return
		}

		fmt.Println(profiles)
		fmt.Println(output)

	case internal.FormatShikari:
		profiles, output, err := shikari.Convert(profiles)
		if err != nil {
			log.Println("error:", err)
			return
		}

		fmt.Println(profiles)
		fmt.Println(output)
	}

}

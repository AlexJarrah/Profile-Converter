package main

import (
	"fmt"
	"log"

	"github.com/quo0001/Profile-Converter/internal"
	"github.com/quo0001/Profile-Converter/internal/cli"
	"github.com/quo0001/Profile-Converter/internal/profiles/shikari"
	"github.com/quo0001/Profile-Converter/internal/profiles/stellar"
)

func main() {
	path, inputFormat, outputFormat := cli.Prompt()

	fmt.Println(path)
	fmt.Println(inputFormat)
	fmt.Println(outputFormat)

	var profiles []internal.Profile
	var err error

	switch inputFormat {
	case internal.FormatStellar:
		profiles, err = stellar.Parse(path)

	case internal.FormatShikari:
		profiles, err = shikari.Parse(path)
	}

	if err != nil {
		log.Println("error parsing input:", err)
		return
	}

	switch outputFormat {
	case internal.FormatStellar:
		profiles, output, err := stellar.Convert(profiles)
		if err != nil {
			log.Println("error converting to Stellar:", err)
			return
		}

		fmt.Println(profiles)
		fmt.Println(output)

	case internal.FormatShikari:
		profiles, output, err := shikari.Convert(profiles)
		if err != nil {
			log.Println("error converting to Shikari:", err)
			return
		}

		fmt.Println(profiles)
		fmt.Println(output)
	}

}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/quo0001/Profile-Converter/internal"
	"github.com/quo0001/Profile-Converter/internal/profiles/stellar"
)

func main() {
	// path, inputFormat, outputFormat := prompt.Prompt()
	path := "formats/profiles/stellar.json"
	inputFormat := internal.FormatStellar
	outputFormat := internal.FormatStellar

	var profiles []internal.Profile
	var err error

	switch inputFormat {
	case internal.FormatStellar:
		if profiles, err = stellar.Parse(path); err != nil {
			fmt.Println("Error:", err)
			return
		}

	case internal.FormatShikari:
	}

	fmt.Println(profiles[0].Billing.Address)

	var output []byte
	switch outputFormat {
	case internal.FormatStellar:
		out := stellar.Convert(profiles)
		if output, err = json.Marshal(out); err != nil {
			fmt.Println("Error:", err)
			return
		}

	case internal.FormatShikari:
	}

	fmt.Println(string(output))
}

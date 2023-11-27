package main

import (
	"fmt"
	"log"

	"github.com/quo0001/Profile-Converter/internal"
	"github.com/quo0001/Profile-Converter/internal/cli"
	"github.com/quo0001/Profile-Converter/internal/profiles/shikari"
	"github.com/quo0001/Profile-Converter/internal/profiles/stellar"
)

func parseInput(path string, inputFormat internal.Format) (profiles []internal.Profile, err error) {
	switch inputFormat {
	case internal.FormatStellar:
		return stellar.Parse(path)
	case internal.FormatShikari:
		return shikari.Parse(path)
	default:
		return nil, fmt.Errorf("invalid input format")
	}
}

func convert(profiles []internal.Profile, outputFormat internal.Format) (any, string, error) {
	switch outputFormat {
	case internal.FormatStellar:
		return stellar.Convert(profiles)
	case internal.FormatShikari:
		return shikari.Convert(profiles)
	default:
		return nil, "", fmt.Errorf("invalid output format")
	}
}

func main() {
	path, inputFormat, outputFormat := cli.Prompt()

	var profiles []internal.Profile
	var err error

	profiles, err = parseInput(path, inputFormat)
	if err != nil {
		log.Println("error parsing input:", err)
		return
	}

	_, output, err := convert(profiles, outputFormat)
	if err != nil {
		log.Println("error converting:", err)
		return
	}

	fmt.Println(output)
}

package main

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Profile-Converter/internal"
	"github.com/AlexJarrah/Profile-Converter/internal/cli"
	"github.com/AlexJarrah/Profile-Converter/internal/profiles/aycd"
	"github.com/AlexJarrah/Profile-Converter/internal/profiles/shikari"
	"github.com/AlexJarrah/Profile-Converter/internal/profiles/stellar"
	"github.com/AlexJarrah/Profile-Converter/internal/profiles/trickle"
	"github.com/atotto/clipboard"
)

func parseInput(path string, inputFormat internal.Format) (profiles []internal.Profile, err error) {
	switch inputFormat {
	case internal.FormatAYCD:
		return aycd.Parse(path)
	case internal.FormatStellar:
		return stellar.Parse(path)
	case internal.FormatShikari:
		return shikari.Parse(path)
	case internal.FormatTrickle:
		return trickle.Parse(path)
	default:
		return nil, fmt.Errorf("invalid input format")
	}
}

func convert(profiles []internal.Profile, outputFormat internal.Format) (any, string, error) {
	switch outputFormat {
	case internal.FormatAYCD:
		return aycd.Convert(profiles)
	case internal.FormatStellar:
		return stellar.Convert(profiles)
	case internal.FormatShikari:
		return shikari.Convert(profiles)
	case internal.FormatTrickle:
		return trickle.Convert(profiles)
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
	clipboard.WriteAll(output)
}

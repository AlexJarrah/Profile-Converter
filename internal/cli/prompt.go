package cli

import (
	"github.com/AlexJarrah/Profile-Converter/internal"
	"github.com/rivo/tview"
)

var options = []string{"AYCD", "Stellar", "Shikari", "Trickle"}

func Prompt() (filePath string, inputFormat internal.Format, outputFormat internal.Format) {
	app := tview.NewApplication()
	setStyles()

	form := tview.NewForm().
		AddInputField("File Path", "", 30, nil, func(text string) {
			filePath = text
		}).
		AddDropDown("Input Format", options, 0, func(option string, optionIndex int) {
			inputFormat = internal.Format(optionIndex)
		}).
		AddDropDown("Output Format", options, 0, func(option string, optionIndex int) {
			outputFormat = internal.Format(optionIndex)
		}).
		AddButton("Generate", func() { app.Stop() }).
		SetButtonsAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return filePath, inputFormat, outputFormat
}

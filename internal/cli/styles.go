package cli

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	MainColor          = tcell.NewRGBColor(26, 26, 26)
	AccentColor        = tcell.NewRGBColor(90, 80, 225)
	BackgroundColor    = tcell.NewRGBColor(10, 10, 10)
	PrimaryTextColor   = tcell.NewRGBColor(250, 250, 250)
	SecondaryTextColor = tcell.NewRGBColor(180, 180, 180)
)

func setStyles() {
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor:    BackgroundColor,
		ContrastBackgroundColor:     MainColor,
		MoreContrastBackgroundColor: BackgroundColor,
		BorderColor:                 AccentColor,
		TitleColor:                  SecondaryTextColor,
		GraphicsColor:               AccentColor,
		PrimaryTextColor:            PrimaryTextColor,
		SecondaryTextColor:          SecondaryTextColor,
		TertiaryTextColor:           SecondaryTextColor,
		InverseTextColor:            BackgroundColor,
		ContrastSecondaryTextColor:  PrimaryTextColor,
	}
}

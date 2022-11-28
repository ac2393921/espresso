package main

import (
	"github.com/ac2393921/tuikan/pkg/gui"
	"github.com/mattn/go-runewidth"
)

func main() {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}

	app := gui.CreateApplication()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
)

func createCommandList() (commandList *tview.List) {
	commandList = tview.NewList()
	commandList.SetBorder(true).SetTitle("Command")
	return commandList
}

func createInfoPanel(app *tview.Application) (infoPanel *tview.Flex) {

	infoTable := tview.NewTable()
	infoTable.SetBorder(true).SetTitle("Information")

	cnt := 0
	infoTable.SetCellSimple(cnt, 0, "Data1:")
	infoTable.GetCell(cnt, 0).SetAlign(tview.AlignRight)
	info1 := tview.NewTableCell("aaa")
	infoTable.SetCell(cnt, 1, info1)
	cnt++

	infoTable.SetCellSimple(cnt, 0, "Data2:")
	infoTable.GetCell(cnt, 0).SetAlign(tview.AlignRight)
	info2 := tview.NewTableCell("bbb")
	infoTable.SetCell(cnt, 1, info2)
	cnt++

	infoTable.SetCellSimple(cnt, 0, "Time:")
	infoTable.GetCell(cnt, 0).SetAlign(tview.AlignRight)
	info3 := tview.NewTableCell("0")
	infoTable.SetCell(cnt, 1, info3)
	cnt++

	infoPanel = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(infoTable, 0, 1, false)

	return infoPanel
}

func createLayout(cList tview.Primitive, recvPanel tview.Primitive) (layout *tview.Flex) {
	bodyLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(cList, 20, 1, true).
		AddItem(recvPanel, 0, 1, false)

	header := tview.NewTextView()
	header.SetBorder(true)
	header.SetText("tview study")
	header.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(bodyLayout, 0, 1, true)

	return layout
}

func createApplication() (app *tview.Application) {
	app = tview.NewApplication()
	pages := tview.NewPages()

	infoPanel := createInfoPanel(app)

	commandList := createCommandList()
	commandList.AddItem("Test", "", 'p', nil)
	commandList.AddItem("Quit", "", 'q', func() {
		app.Stop()
	})
	layout := createLayout(commandList, infoPanel)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true)
	return app
}

func main() {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}

	app := createApplication()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

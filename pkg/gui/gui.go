package gui

import "github.com/rivo/tview"

func createStagePanel(title string) (StagePanel *tview.List) {
	StagePanel = tview.NewList()
	StagePanel.SetBorder(true).SetTitle(title)
	return StagePanel
}

func createLayout(todoPanel, inProgressPanel, donePanel tview.Primitive) (layout *tview.Flex) {
	bodyLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(todoPanel, 0, 1, false).
		AddItem(inProgressPanel, 0, 1, false).
		AddItem(donePanel, 0, 1, false)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(bodyLayout, 0, 1, true)

	return layout
}

func CreateApplication() (app *tview.Application) {
	app = tview.NewApplication()
	pages := tview.NewPages()

	todoPanel := createStagePanel("Todo")
	inProgressPanel := createStagePanel("InProgress")
	donePanel := createStagePanel("Done")

	layout := createLayout(todoPanel, inProgressPanel, donePanel)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true)
	return app
}

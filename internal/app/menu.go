package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MenuItem struct {
	name string
	icon fyne.Resource
	run  func(fyne.Window) fyne.CanvasObject
}

type Menu struct {
	Items []MenuItem
}

func CreateMenu() *Menu {
	return &Menu{
		Items: []MenuItem{
			{
				name: "Record",
				icon: theme.MediaRecordIcon(),
				run:  nil,
			},
			{
				name: "Quit",
				icon: nil,
				run:  nil,
			},
		},
	}
}

func (m *Menu) Show(ui fyne.App) {
	content := container.NewMax()

	w := ui.NewWindow("Examples")

	appList := widget.NewList(
		func() int {
			return len(m.Items)
		},
		func() fyne.CanvasObject {
			icon := &canvas.Image{}
			label := widget.NewLabel("Text Editor")
			labelHeight := label.MinSize().Height
			icon.SetMinSize(fyne.NewSize(labelHeight, labelHeight))
			return container.NewBorder(nil, nil, icon, nil,
				label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			img.Resource = m.Items[id].icon
			img.Refresh()
			text.SetText(m.Items[id].name)
		})
	appList.OnSelected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{m.Items[id].run(w)}
	}

	split := container.NewHSplit(appList, content)
	split.Offset = 0.1

	w.SetContent(split)
	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()
}

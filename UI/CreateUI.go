package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NewUI struct {
	Win fyne.Window
}

func (u *NewUI) MakeUI(w fyne.Window, a fyne.App) fyne.CanvasObject {
	var data = []string{"a", "string", "list"}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	//VBox := container.NewVBox(bt1, bt2, bt3)
	HSplit := container.NewHSplit(list, widget.NewButton("ffff", func() {

	}))
	HSplit.Offset = 0.2
	return container.NewBorder(nil, nil, nil, nil, HSplit)
}

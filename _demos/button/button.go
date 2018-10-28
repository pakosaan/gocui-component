package main

import (
	"github.com/jroimartin/gocui"
	component "github.com/skanehira/gocui-component"
)

func main() {
	gui, err := gocui.NewGui(gocui.Output256)
	gui.Highlight = true

	if err != nil {
		panic(err)
	}
	defer gui.Close()

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	component.NewButton(gui, "Save", 0, 0, 5).
		AddHandler(component.Handlers{gocui.KeyEnter: quit, gocui.KeyTab: changeButton}).
		SetPrimary().
		Draw()

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func changeButton(g *gocui.Gui, v *gocui.View) error {
	switch v.Name() {
	case "Save":
		g.SetCurrentView("Cancel")
	case "Cancel":
		g.SetCurrentView("Save")
	}

	return nil
}

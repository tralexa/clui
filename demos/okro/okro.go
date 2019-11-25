package okro


import (
	ui "github.com/tralexa/clui"
)

type Okro struct {

}

func (o *Okro) Main() {
	mainLoop()
}

func createView() {
	s := "Select "

	dlg := ui.CreateFileSelectDialog(
		s,
		"",
		"",
		true,
		true)
	dlg.OnClose(func() {
		if !dlg.Selected {
			return
		}

		var lb string
		if dlg.Exists {
			lb = "Selected existing"
		} else {
			lb = "Create new"
		}

		lb += dlg.FilePath
	})
}

func mainLoop() {
	// Every application must create a single Composer and
	// call its intialize method
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	createView()

	// start event processing loop - the main core of the library
	ui.MainLoop()
}
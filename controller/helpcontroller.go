package controller

import (
	"fmt"
	"gitlab_ui/view"
)

type helpController struct {}

var HelpController =&helpController{}

func (H *helpController)PopulateInfo(uiApp *view.UiApp)  {
	uiApp.Mui.HelpText.SetText(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		`CTRL+P [yellow]PREVIEW`,
		`[white]ESC [yellow]CANCEL`,
		`[white]CTRL+Q [yellow]QUIT`,
		`[white]CTRL+A [yellow]ADDUSER`,
		`[white]CTRL+N [yellow]USERNAME`,
		`[white]CTRL+F [yellow]SEARCH`,
	        `[white]CTRL+D [yellow]DELETE`))
}

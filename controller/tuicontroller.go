package controller

import (
	"gitlab_ui/view"
)

type appController struct {
	uiApp *view.UiApp
}

var AppController=&appController{}

func(c *appController)InitApp() *view.UiApp{
	c.uiApp=&view.UiApp{}
	c.uiApp.Mui=&view.Mui{}
	c.uiApp.Mui.InitBox()
	c.uiApp.Mui.InitElem()
	c.uiApp.Mui.InitType()
	c.uiApp.InitPage()
	c.uiApp.InitApp()
	return c.uiApp
}

func (c *appController) StartApp()  {
	err := c.uiApp.App.SetRoot(c.uiApp.Pages, true).Run();
	if err != nil{
		panic(err)
	}
}
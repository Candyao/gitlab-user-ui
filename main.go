package main

import (
	"gitlab_ui/config"
	"gitlab_ui/connect"
	"gitlab_ui/controller"
	"gitlab_ui/event"
)

func main()  {
	config.ConfigManager.Init(nil)
	connect.Conn.Init(config.ConfigManager.Url,config.ConfigManager.Token)
	app:=controller.AppController.InitApp()
	controller.HelpController.PopulateInfo(app)
	controller.UserController.InitpopulateUsers(app)
	event.InitAppEvent(app)
	controller.AppController.StartApp()
}

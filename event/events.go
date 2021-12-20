package event

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gitlab_ui/controller"
	"gitlab_ui/operator"
	"gitlab_ui/view"
)

func InitAppEvent(c *view.UiApp) {
	c.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			keyEsc(c)
		case tcell.KeyCtrlA:
			keyCtrlA(c)
		case tcell.KeyCtrlN:
			keyCtrlN(c)
		case tcell.KeyCtrlP:
			keyCtrlP(c)
		case tcell.KeyCtrlF:
			keyCtrlF(c)
		case tcell.KeyCtrlQ:
			keyQ(c)
		}
		return event
	})
	c.UserList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch  event.Key() {
		case tcell.KeyCtrlD:
			keyCtrlD(c)
		}
		return event
	})
}

const (
	main  = 0
	add   = 1
	modal = 2
	delmodal = 3
)

var page int
var hasFilter int

func FormButton(c *view.UiApp) {
	c.AddUserForm.
		AddButton("Save", func() {
			userName := c.AddUserForm.GetFormItemByLabel("USERNAME").(*tview.InputField).GetText()
			name := c.AddUserForm.GetFormItemByLabel("NAME").(*tview.InputField).GetText()
			passwd := c.AddUserForm.GetFormItemByLabel("PASSWORD").(*tview.InputField).GetText()
			repasswd := c.AddUserForm.GetFormItemByLabel("REPASSWORD").(*tview.InputField).GetText()
			err, user := operator.Adduser(&passwd, &name, &userName, &repasswd)
			if err != nil {
				c.InitModal(err.Error(), func() {
					page = modal
					c.Pages.SwitchToPage("add")
					page = add
				})
				c.Pages.SwitchToPage("modal")
			} else if user.ID > 0 {
				c.InitModal(fmt.Sprintf("Create User %s Success", user.Name), func() {
					page = modal
					keyCtrlA(c)
				})
				c.Pages.SwitchToPage("modal")
				page=modal
			}
		}).
		AddButton("Quit", func() {
			c.Pages.SwitchToPage("main")
			controller.UserController.PopulateUsers(c, nil)
			page = main
		}).
		SetButtonBackgroundColor(tcell.ColorSteelBlue).
		SetButtonTextColor(tcell.ColorWhiteSmoke).
		SetButtonsAlign(tview.AlignCenter)
}

func keyCtrlP(c *view.UiApp) {
	if page == main {
		c.App.SetFocus(c.Mui.PreviewText)
	}
	return
}

func keyCtrlN(c *view.UiApp) {
	if page == main {
		c.App.SetFocus(c.Mui.UserList)
	}
	return
}

func keyCtrlF(c *view.UiApp) {
	if page == main {
		if hasFilter == 0 {
			c.Mui.UserFlexBox.AddItem(c.Mui.UserFilter, 0, 1, false)
			c.Mui.UserFilter.SetText("")
			c.App.SetFocus(c.Mui.UserFilter)
			hasFilter = 1
		} else {
			c.Mui.UserFlexBox.RemoveItem(c.Mui.UserFilter)
			c.App.SetFocus(c.Mui.UserFlexBox)
			hasFilter = 0
		}
	}
	return
}

func keyEsc(c *view.UiApp) {
	if page != main {
		c.Pages.SwitchToPage("main")
		controller.UserController.PopulateUsers(c, nil)
		page = main
	} else {
		c.Mui.UserFlexBox.RemoveItem(c.Mui.UserFilter)
		controller.UserController.PopulateUsers(c, nil)
	}
	c.App.SetFocus(c.Mui.UserFlexBox)
}

func keyCtrlA(c *view.UiApp) {
	if page == main || page == modal {
		c.AddUserForm.Clear(true)
		c.InitForm()
		FormButton(c)
		c.Pages.SwitchToPage("add")
		page = add
	}
	return
}

func keyQ(c *view.UiApp) {
	c.App.Stop()
	return
}

func keyCtrlD(c *view.UiApp) {
	if page == main {
		page = delmodal
		index := c.UserList.GetCurrentItem()
		user := controller.UserController.GetUser(index)
		c.InitDmodal(fmt.Sprintf("do you want delete %s", user.Name),func(s string) {
			if s == "YES" {
				_ = operator.DeletedUser(user.ID)
			}
			keyEsc(c)
		})
		c.Pages.SwitchToPage("delmodal")
	}
	return
}

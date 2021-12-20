package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UiApp struct {
	App  *tview.Application
	Pages *tview.Pages
	ViewFlexBox *tview.Flex
	*Mui
}

type Mui struct {
	UserList       *tview.List
	UserFilter     *tview.InputField
	PreviewText    *tview.TextView
	HelpText       *tview.TextView
	UserFlexBox    *tview.Flex
	PreviewFlexBox *tview.Flex
	AddUserForm    *tview.Form
	HelpFlexBox    *tview.Flex
	AddModal       *tview.Modal
	DelModal       *tview.Modal
}

func (a *UiApp)InitApp()  {
	a.App=tview.NewApplication()
	a.Pages=tview.NewPages()
	a.Pages.AddPage("delmodal",a.DelModal,true,false)
	a.Pages.AddPage("modal",a.AddModal,true,false)
	a.Pages.AddPage("add",a.modal(a.AddUserForm,60,15),true,false)
	a.Pages.AddPage("main",a.ViewFlexBox,true,true)
}

func (a *UiApp) InitPage() {
	a.ViewFlexBox = tview.NewFlex()
	a.ViewFlexBox = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.Mui.HelpFlexBox, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(a.Mui.UserFlexBox, 0, 1, true).
			AddItem(a.Mui.PreviewFlexBox, 0, 10, true), 0, 10, true)
}

func (ui *Mui) InitBox() {
	ui.AddModal = tview.NewModal()
	ui.DelModal = tview.NewModal()
	ui.UserFlexBox = tview.NewFlex()
	ui.PreviewFlexBox = tview.NewFlex()
	ui.HelpFlexBox = tview.NewFlex()
	ui.AddUserForm = tview.NewForm()
	ui.UserFlexBox.SetDirection(tview.FlexRow)
	ui.PreviewFlexBox.SetDirection(tview.FlexRow)
	ui.HelpFlexBox.SetDirection(tview.FlexColumn)
}

func (ui *Mui) InitElem() {
	ui.PreviewText = tview.NewTextView()
	ui.HelpText = tview.NewTextView()
	ui.UserList = tview.NewList()
	ui.UserFilter = tview.NewInputField()
}

func (ui *Mui) InitType() {
	ui.UserList.
		ShowSecondaryText(false).
		SetTitle("USERNAME").
	    SetBorder(true).
		SetBorderColor(tcell.ColorSteelBlue)

	ui.UserFilter.
		SetFieldBackgroundColor(tcell.ColorWhite).
		SetFieldTextColor(tcell.ColorBlack)

	ui.PreviewText.
		SetDynamicColors(true).
		SetRegions(true).
		SetScrollable(true).
		SetTitle("PREVIEW").
		SetBorder(true).
		SetBorderColor(tcell.ColorSaddleBrown)

	ui.HelpText.
		SetDynamicColors(true).
		SetRegions(true).
		SetScrollable(false).
		SetTitle("HELP").
		SetBorder(true).

		SetBorderColor(tcell.ColorSaddleBrown)

	ui.HelpFlexBox.AddItem(ui.HelpText, 0, 1, false)
	ui.UserFlexBox.AddItem(ui.UserList, 0, 1, true)
	ui.PreviewFlexBox.AddItem(ui.PreviewText, 0, 10, true)
}

func (ui *Mui)InitForm()  {
	ui.AddUserForm.
		AddInputField("USERNAME","",32,nil,nil).
		AddInputField("NAME","",32,nil,nil).
		AddPasswordField("PASSWORD","",32,0,nil).
		AddPasswordField("REPASSWORD","",32,0,nil).
		SetFieldBackgroundColor(tcell.ColorBlack)
}

func (a *UiApp)InitModal(message string,doneFunc func()){
	a.AddModal.
		ClearButtons().
		SetText(message).
		SetBackgroundColor(tcell.ColorSteelBlue).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			doneFunc()
	})
}

func (a *UiApp)InitDmodal(message string ,doneFunc func(s string))  {
	a.DelModal.
		ClearButtons().
		SetText(message).
		SetBackgroundColor(tcell.ColorSteelBlue).
		AddButtons([]string{"YES","NO"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			doneFunc(buttonLabel)
		})
}

func (a *UiApp)modal(p tview.Primitive,width,height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0,width,0).
		SetRows(0,height,0).
		AddItem(p,1,1,1,1,0,0,true)
}

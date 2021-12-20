package controller

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gdamore/tcell/v2"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/xanzy/go-gitlab"
	"gitlab_ui/operator"
	"gitlab_ui/view"
	"sort"
)

type userController struct {
	Users []*gitlab.User
}

var UserController=&userController{}

func (u *userController)InitpopulateUsers(uiApp *view.UiApp)  {
	u.Users = operator.GetAllUser()
	u.PopulateUsers(uiApp,nil)
	u.setChangeFunc(uiApp)
	u.setFinishedFunc(uiApp)
}

func (u *userController) PopulateUsers(uiApp *view.UiApp,index []int) {
	uiApp.Mui.UserList.Clear()
	u.Users = operator.GetAllUser()
	if index != nil {
		for _,vl:= range index {
			uiApp.Mui.UserList.AddItem(u.Users[vl].Name, "", 0, nil)
		}
	} else {
		for _, v := range u.Users {
			uiApp.Mui.UserList.AddItem(v.Name, "", 0, nil)
		}
	}
	return
}

func (u *userController) setChangeFunc(uiApp *view.UiApp) {
	uiApp.Mui.UserList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		var preview string
		for lk, lv := range u.Users {
			if lv.Name == mainText {
				index = lk
			}
		}
		m := structs.Map(&u.Users[index])
		keys := make([]string, len(m))
		i := 0
		for k, _ := range m {
			keys[i] = k
			i++
		}
		sort.Strings(keys)

		for _, v := range keys {
			preview += fmt.Sprintf("%s: %v\n", v, m[v])
		}
		uiApp.Mui.PreviewText.SetText(preview)
		return
	})
	return
}

func (u *userController) setFinishedFunc(uiApp *view.UiApp) {
	uiApp.Mui.UserFilter.SetFinishedFunc(func(key tcell.Key) {
		var lname []int
		name := uiApp.Mui.UserFilter.GetText()
		for lk, lv := range u.Users {
			if fuzzy.Match(name,lv.Name){
				lname=append(lname, lk)
			}
		}
		u.PopulateUsers(uiApp,lname)
		uiApp.Mui.UserFlexBox.RemoveItem(uiApp.Mui.UserFilter)
		uiApp.App.SetFocus(uiApp.Mui.UserList)
		return
	})
}

func (u *userController) GetUser(index int) *gitlab.User {
	return u.Users[index]
}
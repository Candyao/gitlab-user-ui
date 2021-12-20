package operator

import (
	"github.com/xanzy/go-gitlab"
	"gitlab_ui/connect"
	"log"
)

func GetAllUser()([]*gitlab.User){
	users,_,err:=connect.Conn.Client.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err!=nil{
		log.Fatalf("list user error %v",err)
	}
	return users
}
package operator

import (
	"errors"
	"github.com/xanzy/go-gitlab"
	"gitlab_ui/config"
	"gitlab_ui/connect"
)

func Adduser(passwd, name, username, repasswd *string) (error, *gitlab.User) {
	if *passwd != *repasswd {
		return errors.New("The password is not the same"), nil
	}
	rest := true
	email := *username + "@" + config.ConfigManager.EmailDomain
	u, _, err := connect.Conn.Client.Users.CreateUser(&gitlab.CreateUserOptions{
		Password:      passwd,
		ResetPassword: &rest,
		Username:      username,
		Name:          name,
		Email:         &email,
	})
	if err != nil {
		return err, nil
	}
	return nil, u
}

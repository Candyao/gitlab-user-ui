package operator

import "gitlab_ui/connect"

func DeletedUser(id int) error {
	_,err:=connect.Conn.Client.Users.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
package connect

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
)


type conn struct {
	Client *gitlab.Client
}

var Conn=&conn{}

func (this *conn)Init(url,token string)  {
	this.Client=this.GetConn(url,token)
}

func (this *conn)GetConn(url,token string)(*gitlab.Client){
	git,err:=gitlab.NewClient(token,gitlab.WithBaseURL(fmt.Sprintf("%s/%s",url,"/api/v4")))
	if err !=nil{
		log.Fatalf("Failed to create client: %v", err)
	}
	return git
}
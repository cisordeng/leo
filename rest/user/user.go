package user

import (
	"github.com/cisordeng/beego/xenon"

	bUser "leo/business/user"
)

type User struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(User))
}

func (this *User) Resource() string {
	return "user.user"
}

func (this *User) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"name",
			"password",
			"avatar",
			"type",
		},
	}
}

func (this *User) Put() {
	name := this.GetString("name", "")
	password := this.GetString("password", "")
	avatar := this.GetString("avatar", "")
	t := this.GetString("type", "")

	user := bUser.GetUserByNameInType(name, t)
	if user != nil {
		xenon.RaiseException("rest:name is exist", "用户名已存在")
	}
	user = bUser.NewUser(name, password, avatar, t)
	data := bUser.EncodeUser(user)
	this.ReturnJSON(data)
}

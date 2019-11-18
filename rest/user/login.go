package user

import (
	"github.com/cisordeng/beego/xenon"

	bUser "leo/business/user"
)

type Login struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Login))
}

func (this *Login) Resource() string {
	return "user.login"
}

func (this *Login) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"name",
			"password",
			"type",
		},
	}
}

func (this *Login) Put() {
	name := this.GetString("name", "")
	password := this.GetString("password", "")
	t := this.GetString("type", "")
	sid := bUser.AuthUser(name, password, t)
	if sid != "" {
		user := bUser.GetUserByNameInType(name, t)
		data := bUser.EncodeUser(user)
		data["sid"] = sid
		this.ReturnJSON(data)
	} else {
		xenon.RaiseException("rest:name or password is wrong", "用户名或密码错误")
	}
}

package user

import (
	"github.com/cisordeng/beego/xenon"

	bUser "leo/business/user"
)

type ValidType struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(ValidType))
}

func (this *ValidType) Resource() string {
	return "user.valid_type"
}

func (this *ValidType) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"token",
			"type",
		},
	}
}

func (this *ValidType) Put() {
	t := this.GetString("type", "")
	user := &bUser.User{}
	this.GetUserFromToken(user)
	isValid := bUser.ValidType(user, t)
	this.ReturnJSON(xenon.Map{
		"is_valid": isValid,
	})
}

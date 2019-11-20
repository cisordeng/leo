package user

import (
	"github.com/cisordeng/beego/xenon"

	bUser "leo/business/user"
)

type Users struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Users))
}

func (this *Users) Resource() string {
	return "user.users"
}

func (this *Users) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *Users) Get() {
	page := this.GetPage()
	filters := this.GetFilters()
	orders := this.GetOrders()
	users, pageInfo := bUser.GetPagedUsers(page, filters, orders...)
	data := bUser.EncodeManyUser(users)
	this.ReturnJSON(xenon.Map{
		"users": data,
		"page_info": pageInfo.ToMap(),
	})
}

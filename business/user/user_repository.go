package user

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "leo/model/user"
)

func GetUsers(filters xenon.Map, orderExprs ...string ) []*User {
	o := orm.NewOrm()
	qs := o.QueryTable(&mUser.User{})

	var models []*mUser.User
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	users := make([]*User, 0)
	for _, model := range models {
		users = append(users, InitUserFromModel(model))
	}
	return users
}

func GetPagedUsers(page *xenon.Paginator, filters xenon.Map, orderExprs ...string ) ([]*User, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mUser.User{})

	var models []*mUser.User
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)

	users := make([]*User, 0)
	for _, model := range models {
		users = append(users, InitUserFromModel(model))
	}
	return users, pageInfo
}

func GetUserByNameInType(name string, t string) (user *User)  {
	model := mUser.User{}
	err := orm.NewOrm().QueryTable(&mUser.User{}).Filter(xenon.Map{
		"name": name,
		"type": mUser.STR2USERTYPE[t],
	}).One(&model)
	xenon.PanicNotNilError(err, "raise:user:not_exits", "用户不存在")
	user = InitUserFromModel(&model)
	return user
}

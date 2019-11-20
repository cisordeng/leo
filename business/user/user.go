package user

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "leo/model/user"

)

type User struct {
	Id int
	Name string
	Password string
	Avatar string
	Type string
	CreatedAt time.Time
}

func init() {
}

func InitUserFromModel(model *mUser.User) *User {
	instance := new(User)
	instance.Id = model.Id
	instance.Name = model.Name
	instance.Password = model.Password
	instance.Avatar = model.Avatar
	instance.Type = mUser.USERTYPE2STR[model.Type]
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewUser(name string, password string, avatar string, t string) (user *User) {
	if NameIsExistInType(name, t) {
		xenon.RaiseException("rest:name is exist", "用户名已存在")
	}

	model := mUser.User{
		Name: name,
		Password: xenon.EncodeMD5(password),
		Avatar: avatar,
		Type: mUser.STR2USERTYPE[t],
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitUserFromModel(&model)
}
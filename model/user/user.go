package user

import (
	"time"

	"github.com/cisordeng/beego/orm"
)

const USER_TYPE_ADMIN = 0
const USER_TYPE_DEFAULT = 1

var USERTYPE2STR = map[int]string{
	USER_TYPE_ADMIN: "admin",
	USER_TYPE_DEFAULT: "default",
}
var STR2USERTYPE = map[string]int{
	"admin": USER_TYPE_ADMIN,
	"default": USER_TYPE_DEFAULT,
}

type User struct {
	Id int
	Name string
	Password string
	Avatar string
	Type int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *User) TableName() string {
	return "user_user"
}

func init() {
	orm.RegisterModel(new(User))
}

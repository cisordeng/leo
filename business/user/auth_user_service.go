package user

import (
	"encoding/json"

	"github.com/cisordeng/beego"
	"github.com/cisordeng/beego/xenon"

	mUser "leo/model/user"
)

func AuthUser(name string, password string, t string) string {
	user := GetUserByNameInType(name, t)
	userMap := EncodeUser(user)
	if t == user.Type && user.Password == xenon.EncodeMD5(password) {
		decodedByteToken, err := json.Marshal(userMap)
		xenon.PanicNotNilError(err)
		decodedToken := string(decodedByteToken)

		commonKey := beego.AppConfig.String("api::aesCommonKey")
		sid, err := xenon.EncodeAesWithCommonKey(decodedToken, commonKey)
		xenon.PanicNotNilError(err)
		return sid
	}
	return ""
}

func ValidType(user *User, t string) bool {
	return user.Type == t
}

func NameIsExistInType(name string, t string) bool {
	users := GetUsers(xenon.Map{
		"name": name,
		"type": mUser.STR2USERTYPE[t],
	})
	return len(users) > 0
}
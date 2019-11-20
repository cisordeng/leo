package user

import (
	"github.com/cisordeng/beego/xenon"
)

func EncodeUser(user *User) xenon.Map {
	mapUser := xenon.Map{
		"id": user.Id,
		"name": user.Name,
		"avatar": user.Avatar,
		"type": user.Type,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapUser
}

func EncodeManyUser(articles []*User) []xenon.Map {
	mapUsers := make([]xenon.Map, 0)
	for _, article := range articles {
		mapUsers = append(mapUsers, EncodeUser(article))
	}
	return mapUsers
}
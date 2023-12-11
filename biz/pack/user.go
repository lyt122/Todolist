package pack

import (
	"Todolist/biz/dal/db"
	"Todolist/biz/model/model"
	"strconv"
)

func User(data *db.User) *model.User {
	return &model.User{
		UID:      data.Uid,
		Password: data.Password,
		Username: data.Username,
		CreatAt:  strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdateAt: strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}

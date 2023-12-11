package dal

import (
	"Todolist/biz/dal/db"
	"Todolist/biz/middleware"
)

func Init() {
	db.InitDB()
	middleware.Init()
}

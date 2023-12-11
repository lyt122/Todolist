package pack

import (
	"Todolist/biz/dal/db"
	"Todolist/biz/model/model"
	"strconv"
)

func Task(data *db.Task) *model.Task {
	return &model.Task{
		ID:       data.Id,
		Title:    data.Title,
		Content:  data.Content,
		Status:   data.Status,
		CreatAt:  strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdateAt: strconv.FormatInt(data.UpdatedAt.Unix(), 10),
		StartAt:  strconv.FormatInt(data.StartAt.Unix(), 10),
		EndAt:    strconv.FormatInt(data.EndAt.Unix(), 10),
	}
}

func TaskList(data []*db.Task, total int64) *model.TaskList {
	var taskResp []*model.Task
	for _, v := range data {
		taskResp = append(taskResp, Task(v))
	}
	return &model.TaskList{
		Total: total,
		Tasks: taskResp,
	}
}

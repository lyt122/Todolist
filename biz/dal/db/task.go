package db

import (
	"Todolist/pkg/constants"
	"context"
	"time"
)

func CreateTask(ctx context.Context, title, content string, uid, startAt, endAt int64) (*Task, error) {
	var taskResp *Task
	taskResp = &Task{
		Uid:     uid,
		Title:   title,
		Content: content,
		StartAt: time.Unix(startAt, 0),
		EndAt:   time.Unix(endAt, 0),
	}
	err := DB.WithContext(ctx).Table(constants.TaskTable).Create(taskResp).Error
	if err != nil {
		return nil, err
	}
	return taskResp, nil
}

func UpdateTask(ctx context.Context, id, uid int64, status int16) (err error) {
	if id == 0 {
		err = DB.WithContext(ctx).
			Table(constants.TaskTable).
			Where("uid = ?", uid).
			Update("status", status).
			Error
	} else {
		err = DB.WithContext(ctx).
			Table(constants.TaskTable).
			Where("uid = ?", uid).
			Where("id = ?", id).
			Update("status", status).
			Error
	}
	return
}

func DeleteOneTask(ctx context.Context, uid, id int64) (err error) {
	err = DB.WithContext(ctx).
		Table(constants.TaskTable).
		Delete(&Task{Id: id, Uid: uid}).
		Error
	return
}
func DeleteTasks(ctx context.Context, uid int64, status int16) (err error) {
	if status == 2 {
		err = DB.WithContext(ctx).
			Table(constants.TaskTable).
			Delete(&Task{Uid: uid}).
			Error
	}
	if status == 1 {
		err = DB.WithContext(ctx).
			Table(constants.TaskTable).
			Where("status = ?", 1).
			Delete(&Task{Uid: uid}).
			Error
	}
	if status == 0 {
		err = DB.WithContext(ctx).
			Table(constants.TaskTable).
			Where("status = ?", 0).
			Delete(&Task{Uid: uid}).
			Error
	}
	return
}

func QueryTaskListByStatus(ctx context.Context, pageSize, pageNum, uid int64, status int16) ([]*Task, int64, error) {
	var taskResp []*Task
	var count int64
	err := DB.WithContext(ctx).Table(constants.TaskTable).Where("status = ?", status).
		Where("uid = ?", uid).
		Limit(int(pageSize)).
		Offset(int((pageNum - 1) * pageSize)).
		Order("created_at desc").
		Count(&count).
		Find(&taskResp).
		Error
	if err != nil {
		return nil, -1, err
	}
	return taskResp, count, nil
}
func QueryTaskListByKeyWord(ctx context.Context, pageSize, pageNum, uid int64, keyWord string) ([]*Task, int64, error) {
	var taskResp []*Task
	var count int64
	err := DB.WithContext(ctx).Table(constants.TaskTable).
		Where("uid = ?", uid).
		Where("title like ? or content like ?", keyWord, keyWord).
		Limit(int(pageSize)).
		Offset(int((pageNum - 1) * pageSize)).
		Order("created_at desc").
		Count(&count).
		Find(&taskResp).
		Error
	if err != nil {
		return nil, -1, err
	}
	return taskResp, count, nil
}

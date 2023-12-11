package service

import (
	"Todolist/biz/dal/db"
	"Todolist/biz/model/task"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type TaskService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTaskService(ctx context.Context, c *app.RequestContext) *TaskService {
	return &TaskService{ctx: ctx, c: c}
}

func (s *TaskService) CreateTask(req *task.CreateTaskRequest) (*db.Task, error) {
	return db.CreateTask(s.ctx, req.Title, req.Content, GetUidFormContext(s.c), req.StartAt, req.EndAt)
}

func (s *TaskService) UpdateTask(req *task.UpdateTaskRequest) error {
	return db.UpdateTask(s.ctx, req.ID, GetUidFormContext(s.c), req.Status)
}

func (s *TaskService) DeleteOneTask(req *task.DeleteTaskRequest) error {
	return db.DeleteOneTask(s.ctx, GetUidFormContext(s.c), req.ID)
}

func (s *TaskService) DeleteTaskByStatus(req *task.DeleteTaskByStatusRequest) error {
	return db.DeleteTasks(s.ctx, GetUidFormContext(s.c), req.Status)
}

func (s *TaskService) QueryTaskListByStatus(req *task.QueryTaskListByStatusRequest) ([]*db.Task, int64, error) {
	return db.QueryTaskListByStatus(s.ctx, req.PageSize, req.PageNum, GetUidFormContext(s.c), req.Status)
}

func (s *TaskService) QueryTaskListByKeyword(req *task.QueryTaskListByKeywordRequest) ([]*db.Task, int64, error) {
	return db.QueryTaskListByKeyWord(s.ctx, req.PageSize, req.PageNum, GetUidFormContext(s.c), req.Keyword)
}

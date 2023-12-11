package service

import (
	"Todolist/biz/dal/db"
	"Todolist/biz/model/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

func (s *UserService) Login(req *user.LoginRequest) (*db.User, error) {
	return db.LoginCheck(s.ctx, req.Username, req.Password)
}

func (s *UserService) Register(req *user.RegisterRequest) (*db.User, error) {
	return db.CreateUser(s.ctx, req.Username, req.Password)
}

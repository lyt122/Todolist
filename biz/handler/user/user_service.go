// Code generated by hertz generator.

package user

import (
	"Todolist/biz/middleware"
	user "Todolist/biz/model/user"
	"Todolist/biz/pack"
	"Todolist/biz/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(user.RegisterResponse)
	resp.Resp = pack.BuildBaseReap(nil)
	userResp, err := service.NewUserService(ctx, c).Register(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Data = pack.User(userResp)
	pack.SendResponse(c, resp, consts.StatusOK)
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}

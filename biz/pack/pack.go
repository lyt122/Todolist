package pack

import (
	"Todolist/biz/model/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Base struct {
	Code int16  `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Base Base `json:"base"`
}

func SendResponse(c *app.RequestContext, data interface{}, code int) {
	c.JSON(code, data)
}
func BuildBaseReap(err error) *model.BaseResp {
	if err == nil {
		return &model.BaseResp{
			Code: 10000,
			Msg:  "ok",
		}
	}
	return &model.BaseResp{
		Code: 10001,
		Msg:  err.Error(),
	}
}
func SendFailResponse(c *app.RequestContext, err error) {
	SendResponse(c, BuildBaseReap(err), consts.StatusOK)
}

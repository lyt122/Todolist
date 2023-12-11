package middleware

import (
	"Todolist/biz/dal/db"
	"Todolist/biz/model/user"
	"Todolist/biz/pack"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "uid"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := db.LoginCheck(ctx, loginStruct.Username, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			if users == nil {
				return nil, errors.New("user already exists or wrong password")
			}

			return users, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[IdentityKey]
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Uid,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.SendResponse(c, message, code)
		},
	})
	if err != nil {
		panic(err)
	}
}

func Init() {
	InitJwt()
	errInit := JwtMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}

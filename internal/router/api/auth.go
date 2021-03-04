package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
)

func GetAuth(c *gin.Context)  {
	params := service.AuthRequest{}
	respose := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid{
		global.Logger.Errorf("app.BindAndValid err: %v",errs)
		respose.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	svc := service.New(c)
	err := svc.CheckAuth(&params)
	if err != nil{
		global.Logger.Errorf("app.checkauth err: %v",err)
		respose.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}


	fmt.Println("AppKey",params.AppKey)
	fmt.Println("AppSecret",params.AppSecret)
	token, err := app.GenerateToken(params.AppKey, params.AppSecret)

	if err != nil{
		global.Logger.Errorf("app.GenrateToken err: %v",err)
		respose.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}

	respose.ToResponse(gin.H{
		"token":token,
	})
	return
}

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
)

type Article struct {

}

func NewArticle() Article {
	return Article{}
}

func (a Article)Get(c *gin.Context)  {
	fmt.Println(".......进入了")
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

func (a Article)List(c *gin.Context)  {}

func (a Article)Create(c *gin.Context)  {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid{
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CrateArticle(&param)
	if err != nil{
		global.Logger.Errorf("svc.CreateTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	responseData := response.BuildData(200,"success",nil)
	response.ToResponse(responseData)
}

func (a Article)Update(c *gin.Context)  {}

func (a Article)Delete(c *gin.Context)  {}
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/dao"
	"go-blog/internal/model"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
	"sync"
)

type Article struct {
	ID uint32 `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
	Tag *model.Tag `json:"tag"`
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
	param := dao.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid{
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())

	_,err := svc.CrateArticle(&param)
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


var rwMutex sync.RWMutex

var data sync.Map

func (a Article)Lists(c *gin.Context)  {
	//实例化对应的service
	svc := service.New(c.Request.Context())
	response := app.NewResponse(c)
	lists, err := svc.Lists()
	if err != nil{
		global.Logger.Errorf("svc.listss err: %v",err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	data.Store("list",lists)
	load, ok := data.Load("list")
	if ok{
		response.ToResponse(load)
	}

	return
}


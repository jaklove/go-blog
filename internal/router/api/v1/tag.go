package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
    "go-blog/pkg/convert"
)

type Tag struct {}

func NewTga()Tag  {
	return Tag{}
}


func (t Tag)List(c *gin.Context)  {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid{
		global.Logger.Errorf("app.BindAndValid errs: %v",errors)
		errResp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c),PageSize: app.GetPageSize(c)}
	totalRows,err := svc.CountTag(&service.CountTagRequest{Name: param.Name,State: param.State})
	if err != nil{
		global.Logger.Errorf("svc.CountTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil{
		global.Logger.Errorf("svc.GetTagList err: %v",err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags,totalRows)
	return
}


func (t Tag)Get(c *gin.Context)  {}

func (t Tag)Create(c *gin.Context)  {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid{
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CrateTag(&param)
	if err != nil{
		global.Logger.Errorf("svc.CreateTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

    response.ToResponse(gin.H{})
	return
}

func (t Tag)Update(c *gin.Context)  {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustInt32(),
	}

	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid{
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
    err := svc.UpdateTag(&param)
    if err != nil{
    	global.Logger.Errorf("svc.UpdateTag err: %v",err)
    	response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (t Tag)Delete(c *gin.Context)  {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustInt32(),
	}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid{
		global.Logger.Errorf("app.BindAndValid errs: %v",errors)
		errResp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
    err := svc.DeleteTag(&param)
    if err != nil{
    	global.Logger.Errorf("svc.deleteTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}




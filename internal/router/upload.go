package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/convert"
	"go-blog/pkg/errcode"
	"go-blog/pkg/upload"
)

type Upload struct {}

func NewUpload()Upload  {
	return Upload{}
}

func (u Upload)UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file, header, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	fmt.Println("上传文件")
    if err != nil{
    	errRsp := errcode.InvalidParams.WithDetails(err.Error())
    	response.ToErrorResponse(errRsp)
		return
	}
	if header == nil || fileType <= 0{
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	uploadFile, err := svc.UploadFile(upload.FileType(fileType), file, header)
	if err != nil{
		global.Logger.Errorf("svc.UploadFile err:%v",err)
		errresp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errresp)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url":uploadFile.AccessUlr,
	})
	return
}
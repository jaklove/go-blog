package music

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	service "go-blog/internal/service/student"
	"go-blog/pkg/errcode"
	"go-blog/pkg/app"
)

type Student struct{}

func (s Student) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.NewMusicService(c.Request.Context())
	list, err := svc.GetStudentList()
	if err != nil {
		global.Logger.Errorf("svc.getStudentList err: %v", err)
		fmt.Println("err",err.Error())
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	responseData := response.BuildData(200, "success", list)
	response.ToResponse(responseData)
	return
}

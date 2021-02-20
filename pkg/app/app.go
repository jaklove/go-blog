package app

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/errcode"
	"net/http"
)

type Response struct {
	ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx: ctx}
}

func (r *Response)ToResponse(data interface{})  {
	if data == nil{
		data = gin.H{}
	}
	r.ctx.JSON(http.StatusOK,data)
}

func (r *Response)ToErrorResponse(err *errcode.Error)  {
	response := gin.H{"code":err.Code(),"msg":err.Msg()}
	details := err.Details()
	if len(details) > 0{
		response["details"] = details
	}
	r.ctx.JSON(err.StatusCode(),response)
}

func (r *Response)ToResponseList(list interface{},totalRows int)  {
	r.ctx.JSON(http.StatusOK,gin.H{
		"list":list,
		"pager":Pager{
			Page: GetPage(r.ctx),
			PageSize: GetPageSize(r.ctx),
			TotalRows: totalRows,
		},
	})
}






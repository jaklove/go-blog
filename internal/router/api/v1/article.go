package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (a Article)Create(c *gin.Context)  {}

func (a Article)Update(c *gin.Context)  {}

func (a Article)Delete(c *gin.Context)  {}
package v1

import "github.com/gin-gonic/gin"

type Article struct {

}

func NewArticle()Article  {
	return Article{}
}


func (A Article)Get(c *gin.Context)  {}

func (A Article)List(c *gin.Context)  {}

func (A Article)Create(c *gin.Context)  {}

func (A Article)Update(c *gin.Context)  {}

func (A Article)Delete(c *gin.Context)  {}



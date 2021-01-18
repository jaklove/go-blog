package v1

import "github.com/gin-gonic/gin"

type Tag struct {

}

func NewTag()Tag  {
	return  Tag{}
}

func (T Tag)Get(c *gin.Context)  {}

func (T Tag)List(c *gin.Context)  {}

func (T Tag)Create(c *gin.Context)  {}

func (T Tag)Update(c *gin.Context)  {}

func (T Tag)Delete(c *gin.Context)  {}


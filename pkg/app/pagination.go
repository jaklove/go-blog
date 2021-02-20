package app

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/pkg/convert"
)

func GetPage(c *gin.Context)int  {
	mustInt := convert.StrTo(c.Query("page")).MustInt()
	if mustInt <= 0{
		return 1
	}
	return mustInt
}

func GetPageSize(c *gin.Context) int {
	mustInt := convert.StrTo(c.Query("page_size")).MustInt()
	if mustInt <= 0{
		return global.AppSetting.DefaultPageSize
	}
	return mustInt
}

func GetPageOffset(page,pageSize int)int  {
	result := 0
	if page > 0{
		result = (page - 1)*pageSize
	}
	return result
}
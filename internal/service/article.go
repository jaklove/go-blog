package service

import (
	"go-blog/internal/dao"
	"go-blog/internal/model"
)

type CreateArticleRequest struct {
	Title string `form:"title" binding:"required"`
	Desc string `form:"desc" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func (svc *Service)CrateArticle(param *dao.CreateArticleRequest)(*model.Article,error){
	return svc.dao.CreateArticle(param)
}

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

type ArticleRequest struct {
	 ID uint32 `json:"id"`
	 State uint8 `json:"state"`
}

func (svc *Service)CrateArticle(param *dao.CreateArticleRequest)(*model.Article,error){
	return svc.dao.CreateArticle(param)
}

func (svc *Service)Lists()([]*model.ArticleList,error)  {
	return svc.dao.GetArticleList()
}
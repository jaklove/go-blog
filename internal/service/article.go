package service


type CreateArticleRequest struct {
	Title string `form:"title" binding:"required"`
	Desc string `form:"desc" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func (svc *Service)CrateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title,param.Desc,param.Content)
}

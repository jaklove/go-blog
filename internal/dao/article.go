package dao

import (
	"go-blog/internal/model"
	"go-blog/pkg/app"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id"`
	Title         string `form:"title"`
	Desc          string `form:"desc"`
	Content       string `form:"content"`
	CoverImageUrl string `form:"cover_image_url"`
	CreatedBy string `form:"created_by"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}


func (Dao *Dao) CreateArticle(param *CreateArticleRequest) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}

	return article.Create(Dao.engine)
}

func (Dao *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}

	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(Dao.engine, values)
}

func (Dao *Dao)GetArticle(id uint32,state uint8)(model.Article,error)  {
	article := model.Article{Model:&model.Model{ID: id},State: state}
	return article.Get(Dao.engine)
}

func (Dao *Dao)DeleteArticle(id uint32,state uint8)error {
	article := model.Article{Model:&model.Model{ID: id}}
	return article.Delete(Dao.engine)
}

func (Dao *Dao)CountArticleListByTagID(ID uint32,state uint8)(int,error){
	article := model.Article{State: state}
	return article.CountByTagID(Dao.engine,ID)
}

func (Dao *Dao)GetArticleListByTagID(id uint32,state uint8,page,pagesize int) ([]*model.ArticleRow,error) {
	article := model.Article{State: state}
	return article.ListByTagID(Dao.engine,id,app.GetPageOffset(page,pagesize),pagesize)
}

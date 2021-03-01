package dao

import "go-blog/internal/model"

func (Dao *Dao)CreateArticle(title string,desc string,content string)error {
	article := model.Article{
		Title: title,
		Desc: desc,
		Content: content,
	}
	return article.Create(Dao.engine)
}
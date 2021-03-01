package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagId uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

func (a *Article)TableName()string{
	return "blog_article"
}

func (a *ArticleTag)TableName()string  {
	return "blog_article_tag"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

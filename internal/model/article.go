package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

func (a *ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Updates(values).Where("id = ? AND is_del = ?", a.ID, a.IsDel).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, nil
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pagesize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id as article_id", "ar.title as article_title", "ar.desc as article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id as tag_id", "t.name as tag_name"}...)
	if pageOffset >= 0 && pagesize > 0 {
		db = db.Offset(pageOffset).Limit(pagesize)
	}

	articleTag := ArticleTag{}
	article := Article{}
	rows, err := db.Select(fields).Table(articleTag.TableName()+"as at").Joins("Left join `"+Tag{}.TableName()+"` as t on at.tag_id = t.id").Joins(
		"LEFT JOIN `"+article.TableName()+"` as ar on at.article_id = ar.id").Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}
	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	articleTag := ArticleTag{}
	article := Article{}

	err := db.Table(articleTag.TableName()+" AS at").Joins("LEFT JOIN `"+Tag{}.TableName()+"` as t on at.tag_id = t.id").Joins(
		"LEFT JOIN `"+article.TableName()+"` As ar on at.article_id = ar.id").Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}



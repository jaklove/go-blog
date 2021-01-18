package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-blog/global"
	"go-blog/pkg/setting"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Model struct {
	ID uint32 `grom:"primary_key" json:"id"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings)(*gorm.DB,error)  {
	db, err := gorm.Open(databaseSetting.DbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.Username, databaseSetting.Password,
		databaseSetting.Host, databaseSetting.DBname, databaseSetting.Charset, databaseSetting.ParseTime))
	if err != nil{
		return  nil,err
	}

	if global.ServerSetting.RunMode == "debug"{
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return  db,nil
}

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (a Tag) TableName() string {
	return  "blog_tag"
}

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	ConverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

func (a Article)TableName()string  {
	return "blog_article"
}

type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a ArticleTag)TableName()string  {
	return "blog_article_tag"
}



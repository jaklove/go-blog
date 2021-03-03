package dao

import "go-blog/internal/model"

func (Dao *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleID}
	return articleTag.GetByAID(Dao.engine)
}

func (Dao *Dao) GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagId: tagID}
	return articleTag.ListByTID(Dao.engine)
}

func (Dao *Dao) GetArticleTagListByAIDS(articleIDS []uint32) ([]*model.ArticleTag, error) {
	article := model.ArticleTag{}
	return article.ListByAIDs(Dao.engine, articleIDS)
}

func (Dao *Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{CreatedBy: createdBy},
		ArticleId: articleID,
		TagId:     tagID,
	}
	return articleTag.Create(Dao.engine)
}

func (Dao *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleId: articleID}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_ud":      tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(Dao.engine, values)
}

func (Dao *Dao)DeleteArticleTag(articleID uint32)error  {
	articleTag := model.ArticleTag{ArticleId: articleID}
	return articleTag.DeleteOne(Dao.engine)
}



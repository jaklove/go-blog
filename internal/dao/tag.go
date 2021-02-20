package dao

import (
	"go-blog/internal/model"
	"go-blog/pkg/app"
)

func (Dao *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(Dao.engine)
}

func (Dao *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	size := app.GetPageOffset(page, pageSize)
	return tag.List(Dao.engine, size, pageSize)
}

func (Dao *Dao) CreateTag(name string, state uint8, cratedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: cratedBy},
	}
	return tag.Create(Dao.engine)
}

func (Dao *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if name != "" {
		values["name"] = name
	}
	return tag.Update(Dao.engine, values)
}

func (Dao *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(Dao.engine)
}

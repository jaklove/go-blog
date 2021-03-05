package dao

import "go-blog/internal/model"

func (Dao *Dao)GetStudentList()([]*model.StudentList,error)  {
	userPublicInfo := model.UserPublicInfo{}
	return userPublicInfo.GetList(Dao.engine)
}
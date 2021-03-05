package service

import (
	"context"
	"go-blog/global"
	"go-blog/internal/dao"
	"go-blog/internal/model"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}
func NewMusicService(ctx context.Context)Service  {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.MusicDBEngine)
	return svc
}

func (svc *Service)GetStudentList()([]*model.StudentList,error)  {
	return svc.dao.GetStudentList()
}



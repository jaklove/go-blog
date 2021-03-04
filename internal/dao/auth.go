package dao

import "go-blog/internal/model"

func (Dao *Dao)GetAuth(appKey,appSecret string)(model.Auth,error)  {
	auth := model.Auth{AppKey: appKey,AppSecret: appSecret}
	return auth.Get(Dao.engine)
}



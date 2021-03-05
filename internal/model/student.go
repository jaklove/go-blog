package model

import "github.com/jinzhu/gorm"

type UserPublicInfo struct {
	*Model
}

func (u UserPublicInfo) TableName() string {
	return "user_public_info"
}

type StudentList struct {
	UserId     uint32 `json:"user_id"`
	OpenId     string `json:"open_id"`
	WechatNick string `json:"wechat_nick"`
	KefuId     uint32 `json:"kefu_id"`
}

func (u UserPublicInfo) GetList(db *gorm.DB) ([]*StudentList, error) {
	fields := []string{"user_id", "open_id", "wechat_nick", "kefu_id"}
	tableName := u.TableName()
	rows, err := db.Select(fields).Table(tableName).Limit(10).Rows()
	if err != nil {
		return nil, err
	}

	var studentListLists []*StudentList
	for rows.Next() {
		r := &StudentList{}
		if err := rows.Scan(&r.WechatNick, &r.OpenId, &r.WechatNick, &r.KefuId); err != nil {
			return nil, err
		}

		studentListLists = append(studentListLists, r)
	}

	return studentListLists, nil
}

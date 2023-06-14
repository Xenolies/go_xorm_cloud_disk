package models

import "time"

/**
* @Author: Xenolies
* @Date: 2023/6/14 11:07
* @Version: 1.0
 */

type ShareBasic struct {
	Id                 int
	Identity           string
	RepositoryIdentity string
	ExpiredTime        int
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}

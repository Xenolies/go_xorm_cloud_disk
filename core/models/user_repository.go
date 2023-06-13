package models

import (
	"time"
)

/**
* @Author: Xenolies
* @Date: 2023/6/5 8:56
* @Version: 1.0
 */

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int
	RepositoryIdentity string
	Name               string
	Ext                string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
}

func (u *UserRepository) TableName() string {
	return "user_repository"
}

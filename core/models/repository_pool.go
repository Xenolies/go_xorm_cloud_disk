package models

import (
	"time"
)

/**
* @Author: Xenolies
* @Date: 2023/6/5 8:52
* @Version: 1.0
 */

type RepositoryPool struct {
	Id        int
	Identity  string
	Hash      string
	Name      string
	Ext       string
	Size      int64
	Path      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (r *RepositoryPool) TableName() string {
	return "repository_pool"
}

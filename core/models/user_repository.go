package models

import (
	"time"
)

/**
* @Author: Xenolies
* @Date: 2023/6/5 8:56
* @Version: 1.0
 */

// UserRepository 个人存储池
type UserRepository struct {
	Id                 int       // 个人存储池 ID
	Identity           string    // 个人存储池唯一标识
	UserIdentity       string    // 用户的唯一标识
	ParentId           int       // 用户上传父级 ID
	RepositoryIdentity string    // 关联中心存储池的标识
	Name               string    // 文件名称
	Ext                string    // 扩展名 (显示文件夹类型等信息)
	CreatedAt          time.Time `xorm:"created"` // 文件创建时间
	UpdatedAt          time.Time `xorm:"updated"` // 文件更新时间
	DeletedAt          time.Time `xorm:"deleted"` // 文件删除时间
}

func (u *UserRepository) TableName() string {
	return "user_repository"
}

package models

import (
	"time"
)

/**
* @Author: Xenolies
* @Date: 2023/6/5 8:52
* @Version: 1.0
 */

// RepositoryPool 存储池
type RepositoryPool struct {
	Id        int       // 文件表 ID
	Identity  string    // 文件的唯一标识
	Hash      string    // 文件的哈希值
	Name      string    // 文件名称
	Ext       string    // 扩展名
	Size      int64     // 文件大小
	Path      string    // 文件路径
	CreatedAt time.Time `xorm:"created"` // 文件创建时间
	UpdatedAt time.Time `xorm:"updated"` // 文件更新时间
	DeletedAt time.Time `xorm:"deleted"` // 文件删除时间
}

func (r *RepositoryPool) TableName() string {
	return "repository_pool"
}

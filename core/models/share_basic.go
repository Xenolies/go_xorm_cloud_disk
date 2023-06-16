package models

import "time"

/**
* @Author: Xenolies
* @Date: 2023/6/14 11:07
* @Version: 1.0
 */

// ShareBasic 文件分享表
type ShareBasic struct {
	Id                     int       // 分享表 ID
	Identity               string    // 公共池中唯一标识
	UserIdentity           string    // 分享用户的唯一标识
	RepositoryIdentity     string    // 公共池中的唯一标识
	UserRepositoryIdentity string    // 用户池中的唯一标识
	ExpiredTime            int       // 失效时间，单位秒, [0- 永不失效]
	ClickNum               int       // 点击次数
	CreatedAt              time.Time `xorm:"created"` // 分享创建时间
	UpdatedAt              time.Time `xorm:"updated"` // 分享更新时间
	DeletedAt              time.Time `xorm:"deleted"` // 分享删除时间
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}

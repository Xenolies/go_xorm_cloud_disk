package models

import "time"

/**
* @Author: Xenolies
* @Date: 2023/6/14 11:07
* @Version: 1.0
 */

// UserBasic 用户表
type UserBasic struct {
	Id        int       // 用户表 ID
	Identity  string    // 用户唯一标识
	Name      string    // 用户登录名称
	Password  string    // 用户密码 (取 MD5)
	Email     string    // 注册邮箱
	CreatedAt time.Time `xorm:"created"` // 用户创建时间
	UpdatedAt time.Time `xorm:"updated"` // 用户更新时间
	DeletedAt time.Time `xorm:"deleted"` // 用户删除时间
}

func (ub UserBasic) TableName() string {
	return "user_basic"
}

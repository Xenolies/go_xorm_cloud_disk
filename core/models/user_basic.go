package models

type UserBasic struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (ub UserBasic) TableName() string {
	return "user_basic"
}

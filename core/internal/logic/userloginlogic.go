package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/helper"
	"context"
	"errors"
	"fmt"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	//1. 从数据库中读取数据
	user := new(models.UserBasic)
	get, err := models.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		fmt.Println("models.Engine.Where")
		return nil, err
	}
	if !get {
		fmt.Println(err)
		return nil, errors.New(" 用户名或者密码错误 ")
	}
	// 2. 生成 Token
	token, err := helper.MakeToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	return
}
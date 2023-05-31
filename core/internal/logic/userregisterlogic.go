package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/helper"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// 判断验证码 Code 是否一致
	code, err := models.RDB.Get(l.ctx, req.Email).Result() // 从 redis 中获取验证码
	if err != nil {
		return nil, err
	}
	fmt.Println("REDIS CODE: ", code)
	fmt.Println("req.CODE: ", req.Code)
	if code != req.Code {
		// 如果验证码不相同
		err = errors.New(" 验证码错误!!")
		return
	}
	// 判断用户名是否存在
	cnt, err := models.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}

	if cnt > 0 {
		err = errors.New("用户名存在")
		return
	}
	// 不存在就数据入库
	user := &models.UserBasic{
		Identity:  helper.GetUUID(),
		Name:      req.Name,
		Password:  helper.Md5(req.Password),
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}
	fmt.Println(user.Name)
	fmt.Println(user.Identity)
	insert, err := models.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("INSERT ROW: ", insert)
	return
}

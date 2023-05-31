package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/models"
	"cloud_disk/helper"
	"context"
	"errors"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeSendLogic {
	return &EmailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeSendLogic) EmailCodeSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	// 查询邮箱是否被注册
	cnt, err := models.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, errors.New("邮箱验证码不存在")
	}
	if cnt > 0 {
		err = errors.New(" 该邮箱已经被注册了 ")
		return
	}
	// 不存在则会生成验证码
	code := helper.MakeRandCode()
	// 将生成的验证码发送到邮箱
	models.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpiration))
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	resp = new(types.MailCodeSendReply)
	resp.Code = code
	return
}

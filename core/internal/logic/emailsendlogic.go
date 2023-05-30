package logic

import (
	"cloud_disk/helper"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailSendLogic {
	return &EmailSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailSendLogic) EmailSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	resp = &types.MailCodeSendReply{}
	err = helper.MailSendCode(req.Email, "123123")
	if err != nil {
		return nil, err
	}
	resp.Code = "123123"

	return
}

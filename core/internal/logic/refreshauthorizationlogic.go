package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/helper"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationReply, err error) {
	uc, err := helper.AnalyzeToken(authorization)
	if err != nil {
		return nil, err
	}
	// 重新生成 Token
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.ExpirationTime)
	if err != nil {
		return nil, err
	}
	// 重新生成 用于刷新 Token 的 Token
	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.RefreshAuthorizationReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}

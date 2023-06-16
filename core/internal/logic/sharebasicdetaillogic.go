package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	// 分享记录点击次数更新
	_, err = l.svcCtx.Engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ? ", req.Identity)
	if err != nil {
		return nil, err
	}

	// 获取资源详情
	resp = new(types.ShareBasicDetailReply)
	_, err = l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity , user_repository.name , repository_pool.ext , repository_pool.size , repository_pool.path , repository_pool.hash").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ? ", req.Identity).Get(resp)
	if err != nil {
		return nil, err
	}

	return
}

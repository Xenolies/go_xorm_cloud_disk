package logic

import (
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, uerIdentity string) (resp *types.UserFileMoveReply, err error) {
	// 查询出 ParentID
	parenData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ? ", req.ParentIdentity, uerIdentity).Get(parenData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New(" 该文件夹不存在 ")
	}
	// 更新记录的 ParentID
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(models.UserRepository{
		ParentId: parenData.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

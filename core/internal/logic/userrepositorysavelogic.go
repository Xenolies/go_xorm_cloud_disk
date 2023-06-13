package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/helper"
	"context"
	"fmt"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, UserIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	urp := &models.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       UserIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               req.Name,
		Ext:                req.Ext,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		DeletedAt:          time.Now(),
	}

	_, err = l.svcCtx.Engine.Insert(urp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &types.UserRepositorySaveReply{
		Identity: urp.Identity,
	}, nil
}

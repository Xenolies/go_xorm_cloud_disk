package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/models"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListReply)

	size := req.Size
	if size == 0 {
		size = define.PageSize
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * size
	// 查询用户文件列表
	// 进行链表查询 同时进行分页
	err = l.svcCtx.Engine.Table("user_repository_pool").
		Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).
		Select("user_repository_pool.id, user_repository_pool.repository_identity, user_repository_pool.user_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "user_repository_pool", "user_repository_pool.repository_identity = repository_pool.identity").
		Limit(size, offset).Find(&uf)
	if err != nil {
		return nil, err
	}
	// 计算总页数
	cnt, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).Count(new(models.UserRepositoryPool))
	if err != nil {
		return nil, err
	}
	resp.List = uf
	resp.Count = cnt

	return
}

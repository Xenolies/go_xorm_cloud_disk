package logic

import (
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/helper"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilesUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilesUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilesUploadLogic {
	return &FilesUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilesUploadLogic) FilesUpload(req *types.FilesUploadRequest) (resp *types.FilesUploadReply, err error) {
	rp := &models.RepositoryPool{
		Identity:  helper.GetUUID(),
		Hash:      req.Hash,
		Name:      req.Name,
		Ext:       req.Ext,
		Size:      req.Size,
		Path:      req.Path,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}

	resp = new(types.FilesUploadReply)
	resp.Identity = rp.Identity
	resp.Name = rp.Name
	resp.Ext = rp.Ext
	return
}

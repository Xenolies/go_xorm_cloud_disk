package handler

import (
	"cloud_disk/core/models"
	"cloud_disk/helper"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud_disk/core/internal/logic"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FilesUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilesUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 判断文件是否存在
		file, fileHandler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		b := make([]byte, fileHandler.Size)
		_, err = file.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}

		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FilesUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}
		// 将文件上传到阿里云存储
		uploadPath, err := helper.AliOssUpload(svcCtx.AliOSS, r)
		if err != nil {
			return
		}
		// 向 Logic 中传递文件信息
		req.Name = fileHandler.Filename
		req.Size = fileHandler.Size
		req.Ext = path.Ext(fileHandler.Filename)
		req.Hash = hash
		domain := svcCtx.Config.AliOSS.CustomDomain
		fmt.Println("len(domain): ", len(domain))
		if len(domain) > 0 {
			req.Path = domain + "/" + uploadPath
		} else {
			req.Path = svcCtx.Config.AliOSS.BucketName + "." + svcCtx.Config.AliOSS.Endpoint + "/" + uploadPath
		}

		l := logic.NewFilesUploadLogic(r.Context(), svcCtx)
		resp, err := l.FilesUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

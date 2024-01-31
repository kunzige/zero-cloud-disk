package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/app/files/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListByPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListByPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListByPathLogic {
	return &FileListByPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListByPathLogic) FileListByPath(req *types.FileListByPathRequest) (resp *types.FileListByPathResponse, err error) {
	// todo: add your logic here and delete this line
	userFiles, err := l.svcCtx.TbUserFileModel.FindListByPath(l.ctx, req.UserEmail, req.Path)
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("获取用户目录文件列表失败: %w", err)
	}
	data := []map[string]string{}
	for i := 0; i < len(userFiles); i++ {
		data = append(data, map[string]string{
			"file_name":   userFiles[i].FileName,
			"file_size":   fmt.Sprintf("%d", userFiles[i].FileSize),
			"file_addr":   userFiles[i].FileAddr,
			"upload_at":   userFiles[i].UploadAt.String(),
			"last_update": userFiles[i].LastUpdate.String(),
		})
	}
	if err == models.ErrNotFound {
		return &types.FileListByPathResponse{}, nil
	}

	return &types.FileListByPathResponse{Data: data}, nil
}

package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/app/user/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.FileListRequest) (resp *types.FileListResponse, err error) {
	// todo: add your logic here and delete this line
	userFiles, err := l.svcCtx.TbUserFileModel.FindListByEmail(l.ctx, req.UserEmail)
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("获取用户文件列表失败: %w", err)
	}
	data := []map[string]string{}
	for i := 0; i < len(userFiles); i++ {
		data = append(data, map[string]string{
			"file_name":   userFiles[i].FileName,
			"file_size":   fmt.Sprintf("%d", userFiles[i].FileSize),
			"upload_at":   userFiles[i].UploadAt.String(),
			"last_update": userFiles[i].LastUpdate.String(),
		})
	}
	if err == models.ErrNotFound {
		return &types.FileListResponse{}, nil
	}

	return &types.FileListResponse{Data: data}, nil
}

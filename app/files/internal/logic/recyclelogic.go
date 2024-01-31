package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/app/files/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecycleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecycleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecycleLogic {
	return &RecycleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecycleLogic) Recycle(req *types.FileRecycleListRequest) (resp *types.FileRecycleListResponse, err error) {
	// todo: add your logic here and delete this line
	deletedFiles, err := l.svcCtx.TbUserFileModel.FindListByRecycle(l.ctx, req.UserEmail)
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("获取用户文件列表失败: %w", err)
	}
	data := []map[string]string{}
	for i := 0; i < len(deletedFiles); i++ {
		data = append(data, map[string]string{
			"file_name":   deletedFiles[i].FileName,
			"file_size":   fmt.Sprintf("%d", deletedFiles[i].FileSize),
			"upload_at":   deletedFiles[i].UploadAt.String(),
			"last_update": deletedFiles[i].LastUpdate.String(),
		})
	}
	if err == models.ErrNotFound {
		return &types.FileRecycleListResponse{}, nil
	}

	return &types.FileRecycleListResponse{Data: data}, nil
}

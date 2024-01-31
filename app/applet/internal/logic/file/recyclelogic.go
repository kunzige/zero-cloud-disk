package file

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"
	"zero-cloud-disk/app/user/models"

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

	deletedFiles, err := l.svcCtx.FileRpcClient.FileRecycleList(l.ctx, &pb.FileRecycleListReq{UserEmail: req.UserEmail})
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("获取用户文件列表失败: %w", err)
	}
	data := []map[string]string{}
	for i := 0; i < len(deletedFiles.RecycleFiles); i++ {
		data = append(data, map[string]string{
			"file_name":   deletedFiles.RecycleFiles[i].FileName,
			"file_size":   fmt.Sprintf("%d", deletedFiles.RecycleFiles[i].FileSize),
			"upload_at":   deletedFiles.RecycleFiles[i].UploadAt,
			"last_update": deletedFiles.RecycleFiles[i].LastUpdate,
		})
	}
	if err == models.ErrNotFound {
		return &types.FileRecycleListResponse{}, nil
	}

	return &types.FileRecycleListResponse{Data: data}, nil
}

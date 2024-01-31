package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/models"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileRecycleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileRecycleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileRecycleListLogic {
	return &FileRecycleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileRecycleListLogic) FileRecycleList(in *pb.FileRecycleListReq) (*pb.FileRecycleListRes, error) {
	// todo: add your logic here and delete this line
	deletedFiles, err := l.svcCtx.TbUserFileModel.FindListByRecycle(l.ctx, in.UserEmail)
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("获取用户文件列表失败: %w", err)
	}
	data := make([]*pb.RecycleFile, 0, len(deletedFiles))
	for i := 0; i < len(deletedFiles); i++ {
		data = append(data, &pb.RecycleFile{
			FileName:   deletedFiles[i].FileName,
			FileSize:   deletedFiles[i].FileSize,
			UploadAt:   deletedFiles[i].UploadAt.String(),
			LastUpdate: deletedFiles[i].LastUpdate.String(),
		})
	}
	if err == models.ErrNotFound {
		return &pb.FileRecycleListRes{}, nil
	}

	return &pb.FileRecycleListRes{RecycleFiles: data}, nil
}

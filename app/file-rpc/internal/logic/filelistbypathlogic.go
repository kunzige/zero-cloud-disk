package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/models"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListByPathLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileListByPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListByPathLogic {
	return &FileListByPathLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileListByPathLogic) FileListByPath(in *pb.FileListByPathReq) (*pb.FileListByPathRes, error) {
	// todo: add your logic here and delete this line
	userFiles, err := l.svcCtx.TbUserFileModel.FindListByPath(l.ctx, in.UserEmail, in.Path)
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("获取用户目录文件列表失败: %w", err)
	}
	if err == models.ErrNotFound {
		return &pb.FileListByPathRes{}, nil
	}

	data := make([]*pb.File, 0, len(userFiles))
	for _, file := range userFiles {
		data = append(data, &pb.File{
			FileName:   file.FileName,
			FileSize:   fmt.Sprintf("%d", file.FileSize),
			FileAddr:   file.FileAddr,
			UploadAt:   file.UploadAt.String(),
			LastUpdate: file.LastUpdate.String(),
		})
	}
	return &pb.FileListByPathRes{Files: data}, nil
}

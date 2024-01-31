package file

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"

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
	fileListByPathRes, err := l.svcCtx.FileRpcClient.FileListByPath(l.ctx, &pb.FileListByPathReq{
		UserEmail: req.UserEmail,
		Path:      req.Path,
	})

	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.FileListByPathResponse{Data: fileListByPathRes}, nil
}

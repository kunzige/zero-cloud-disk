package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDeleteLogic) FileDelete(req *types.FileDeleteRequest) (resp *types.FileDeleteResponse, err error) {
	// todo: add your logic here and delete this line
	// 删除文件
	err = l.svcCtx.TbUserFileModel.DeleteByHash(l.ctx, req.FileHash)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &types.FileDeleteResponse{Message: "删除成功"}, nil
}

package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileModifyLogic {
	return &FileModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileModifyLogic) FileModify(req *types.FileModifyRequest) (resp *types.FileModifyResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.TbFileModel.UpdateBySha1(l.ctx, req.FileHash, req.NewFileName)
	if err != nil {
		return nil, fmt.Errorf("更新文件信息失败")
	}
	return &types.FileModifyResponse{Message: "ok"}, nil
}

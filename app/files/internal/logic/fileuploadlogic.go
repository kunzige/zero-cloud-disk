package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/app/files/models"
	"zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Files  []*multipart.FileHeader
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		Files:  []*multipart.FileHeader{},
	}
}

func (l *FileUploadLogic) FileUpload() (resp *types.FileUploadResponse, err error) {
	// todo: add your logic here and delete this line

	// 保存文件
	file, err := l.Files[0].Open()
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}
	os.MkdirAll("./public/files/", os.ModePerm)
	save, _ := os.Create("./public/files/" + l.Files[0].Filename)
	defer save.Close()

	io.Copy(save, file)

	// 上传完文件之后更新数据库的信息
	// 先获取文件的源信息

	fileMeta := models.TbFile{
		FileName: l.Files[0].Filename,
		FileAddr: "/public/files/",
		FileSize: l.Files[0].Size,
		Status:   1,
	}
	message, err := json.Marshal(fileMeta)
	if err != nil {
		panic(err.Error())
	}

	encrypt_hash := utils.Encrypt(string(message))

	fileMeta.FileSha1 = encrypt_hash

	_, err = l.svcCtx.TbFileModel.Insert(l.ctx, &fileMeta)

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("上传文件失败")
	}

	return &types.FileUploadResponse{Message: "file upload ok!"}, nil
}

package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

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

func (l *FileUploadLogic) FileUpload(req types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	// todo: add your logic here and delete this line

	// 保存文件
	file, err := l.Files[0].Open()
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("上传文件失败")
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
	}
	message, err := json.Marshal(fileMeta)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("上传文件失败")
	}

	encrypt_hash := utils.Encrypt(string(message))

	fileMeta.FileSha1 = encrypt_hash

	// 用户与文件关联
	userFile := models.TbUserFile{UserName: req.UserName, FileName: l.Files[0].Filename,
		UserEmail: req.UserEmail, FileSize: l.Files[0].Size, FileSha1: encrypt_hash,
		UploadAt: time.Now(), LastUpdate: time.Now(), Status: 0,
	}

	// _, err = l.svcCtx.TbFileModel.Insert(l.ctx, &fileMeta)
	err = l.svcCtx.TbFileModel.Upload(&fileMeta, &userFile)

	if err != nil {
		return nil, err
	}

	return &types.FileUploadResponse{Message: "file upload ok!"}, nil
}

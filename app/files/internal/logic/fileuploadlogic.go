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
	for i := 0; i < len(l.Files); i++ {

		file, err := l.Files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("上传文件失败")
		}

		if l.Files[i].Size < 2<<20 {
			// 第一种方式
			os.MkdirAll("./public/files/", os.ModePerm)
			save, _ := os.Create("./public/files/" + l.Files[i].Filename)
			defer save.Close()
			io.Copy(save, file)
		} else {
			// 第二种方式
			save, _ := os.OpenFile("./public/files/"+l.Files[i].Filename, os.O_RDWR|os.O_CREATE, 0600)
			defer save.Close()
			buf := make([]byte, 1<<20)
			for {
				n, err := file.Read(buf)
				if err != nil && err != io.EOF {
					fmt.Println(err)
					return nil, fmt.Errorf("上传文件失败")
				}
				if n == 0 {
					break
				}
				save.Write(buf)
			}
		}

		// // 分片处理试试
		// chunks := 1 << 20
		// block := int(math.Ceil(float64(l.Files[i].Size / int64(chunks))))
		// index := 0
		// // 缓冲区
		// var buf []byte
		// for {
		// 	if int64(chunks) > (l.Files[i].Size - int64(index*chunks)) {
		// 		// 获取分片，取余量与chunks的较小值
		// 		buf = make([]byte, l.Files[i].Size-int64(index*chunks))
		// 	} else {
		// 		// make的话长度不能太长 1<<60，或者负数
		// 		buf = make([]byte, chunks)
		// 	}
		// 	_, err := file.Read(buf)
		// 	if err != nil && err != io.EOF {
		// 		fmt.Println(err)
		// 		return nil, fmt.Errorf("分片文件失败")
		// 	}
		// 	saveBlock(fmt.Sprintf("yzk"+"_%d"+".mp4", index), buf)
		// 	if index == block {
		// 		break
		// 	}

		// 	index++
		// }

		// 上传完文件之后更新数据库的信息
		// 先获取文件的源信息

		fileMeta := models.TbFile{
			FileName: l.Files[i].Filename,
			FileAddr: "/public/files/",
			FileSize: l.Files[i].Size,
		}
		message, err := json.Marshal(fileMeta)
		if err != nil {
			fmt.Println(err.Error())
			return nil, fmt.Errorf("上传文件失败")
		}

		encrypt_hash := utils.Encrypt(string(message))

		fileMeta.FileSha1 = encrypt_hash

		// 用户与文件关联
		userFile := models.TbUserFile{UserName: req.UserName, FileName: l.Files[i].Filename,
			UserEmail: req.UserEmail, FileSize: l.Files[i].Size, FileSha1: encrypt_hash,
			UploadAt: time.Now(), LastUpdate: time.Now(), Status: 0,
		}

		err = l.svcCtx.TbFileModel.Upload(&fileMeta, &userFile)

		if err != nil {
			return nil, err
		}

	}
	return &types.FileUploadResponse{Message: "file upload ok!"}, nil

}

func saveBlock(fileName string, buf []byte) {
	save, _ := os.OpenFile("./public/files/"+fileName, os.O_RDWR|os.O_CREATE, 0600)
	defer save.Close()
	save.Write(buf)
}

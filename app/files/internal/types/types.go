// Code generated by goctl. DO NOT EDIT.
package types

type FileDownloadRequest struct {
	FileHash string `form:"file_hash"`
}

type FileDownloadResponse struct {
	Message  string `json:"message"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileSize int64  `json:"file_size"`
}

type FileInfoRequest struct {
	FileHash string `form:"file_hash"`
}

type FileInfoResponse struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type FileModifyRequest struct {
	FileHash    string `form:"file_hash"`
	NewFileName string `form:"new_file_name"`
}

type FileModifyResponse struct {
	Message string `json:"message"`
}

type FileUploadRequest struct {
	UserName  string `form:"user_name"`
	UserEmail string `form:"user_email"`
}

type FileUploadResponse struct {
	Message string `json:"message"`
}

type FileDeleteRequest struct {
	FileHash string `form:"file_hash"`
}

type FileDeleteResponse struct {
	Message string `json:"message"`
}

type FileListRequest struct {
	UserEmail string `form:"user_email"`
}

type FileListResponse struct {
	Data interface{} `json:"data"`
}

type FileMoveRequest struct {
	FileHash string `form:"file_hash"`
	NewPath  string `form:"new_path"`
}

type FileMoveResponse struct {
	Message string `json:"message"`
}

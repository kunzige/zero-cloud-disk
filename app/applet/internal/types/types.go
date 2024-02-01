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

type CancelFileShareRequest struct {
	ShareUuid string `form:"share_uuid"`
}

type CancelFileShareResponse struct {
	Message string `json:"message"`
}

type CreateFolderRequest struct {
	UserEmail string `form:"user_email"`
	UserName  string `form:"user_name"`
	PathName  string `form:"path_name"`
	ParentDir string `form:"parent_dir"`
}

type CreateFolderResponse struct {
	Message string `json:"message"`
}

type DeleteAllForeverRequest struct {
	UserEmail string `form:"user_email"`
}

type DeleteAllForeverResponse struct {
	Message string `json:"message"`
}

type DeleteForeverRequest struct {
	Hash string `form:"hash"`
}

type DeleteForeverResponse struct {
	Message string `json:"message"`
}

type FileDeleteRequest struct {
	FileHash string `form:"file_hash"`
}

type FileDeleteResponse struct {
	Message string `json:"message"`
}

type FileListByPathRequest struct {
	UserEmail string `form:"user_email"`
	Path      string `form:"path"`
}

type FileListByPathResponse struct {
	Data interface{} `json:"data"`
}

type FileMoveRequest struct {
	FileHash string `form:"file_hash"`
	NewPath  string `form:"new_path"`
}

type FileMoveResponse struct {
	Message string `json:"message"`
}

type FileRecycleListRequest struct {
	UserEmail string `form:"user_email"`
}

type FileRecycleListResponse struct {
	Data interface{} `json:"data"`
}

type FileShareRequest struct {
	FileHash  string `form:"file_hash,optional"`
	UserEmail string `form:"user_email"`
	UserName  string `form:"user_name"`
	Path      string `form:"path,optional"`
	ShareUuid string `form:"share_uuid"`
	Duration  int    `form:"duration"`
}

type FileShareResponse struct {
	Message string `json:"message"`
}

type GetShareRequest struct {
	ShareUuid string `form:"share_uuid"`
}

type GetShareResponse struct {
	Data interface{} `json:"data"`
}

type SaveFileRequest struct {
	ShareIdentity string `form:"share_identity"`
	UserEmail     string `form:"user_email"`
	FileName      string `form:"file_name"`
	FileSize      int64  `form:"file_size"`
	UserName      string `form:"user_name"`
	Type          string `form:"type"`
}

type SaveFileResponse struct {
	Message string `json:"message"`
}

type UserEmailRequest struct {
	Email string `form:"email"`
}

type UserEmailResponse struct {
	Code   string `json:"code"`
	Status int    `json:"status,default=1"`
}

type UserInfoRequest struct {
	UserEmail string `form:"user_email"`
}

type UserInfoResponse struct {
	Data interface{} `json:"data"`
}

type UserLoginRequest struct {
	UserEmail    string `form:"user_email"`
	UserPassword string `form:"user_password"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserRegisterRequest struct {
	UserName     string `form:"user_name"`
	UserPassword string `form:"user_password"`
	Email        string `form:"email"`
	Code         string `form:"code"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}

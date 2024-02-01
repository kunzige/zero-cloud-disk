# zero-cloud-disk
> 基于go-zero,sqlx的在线网盘系统
> 使用goctl作为编译项目和管理服务的工具
> 使用gRPC来构建和处理服务间的请求和响应。
> 使用JWT进行权限校验。

## 启动服务
    cd app/applet
    go run diskservice.go

## 启动微服务
### 文件服务
    cd app/file-rpc/
    go run file.go 

### 用户服务
    cd app/user-rpc/
    go run user.go 

### 分享服务
    cd app/share-rpc/
    go run share.go 


## 文件模块
    文件上传
    小文件直接上传，大文件切片上传
    文件下载
    获取文件信息
    文件删除
    获取文件列表
    移动文件

## 用户模块
    用户注册
    用户登录
    获取用户文件列表

## 分享模块
    分享文件
    取消分享
    获取分享文件列表
    转存文件


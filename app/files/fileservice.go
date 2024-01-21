package main

import (
	"flag"
	"fmt"

	"zero-cloud-disk/app/files/internal/config"
	"zero-cloud-disk/app/files/internal/handler"
	"zero-cloud-disk/app/files/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/fileservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 更改上传文件的最大限制
	c.RestConf.MaxBytes = 1024 << 20
	c.RestConf.Timeout = 60000
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

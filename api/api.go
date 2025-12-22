package main

import (
	"flag"
	"fmt"

	"idrm/api/internal/config"
	"idrm/api/internal/handler"
	"idrm/api/internal/svc"
	"idrm/pkg/validator"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化验证器
	validator.Init()
	fmt.Println("Validator initialized successfully")

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting API server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

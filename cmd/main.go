package main

import (
	"log"
	"github.com/dockerps13/sspanel-metron-go/config"
	"github.com/dockerps13/sspanel-metron-go/routers"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	if err := config.InitDB(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	r := routers.SetupRouter()
	addr := config.Conf.Server.Address
	log.Printf("启动服务，监听地址 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("启动失败: %v", err)
	}
}

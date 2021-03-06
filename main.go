package main

import (
	"flag"
	"github.com/SSunSShine/QAsystem/conf"
	"github.com/SSunSShine/QAsystem/route"
)

func main() {
	// 初始化 配置信息，数据库。。。
	var shouldInit = flag.Bool("init", false, "initialize all")
	flag.Parse()

	if *shouldInit {
		initAll(conf.Config())
	}

	route.InitRouter()
}

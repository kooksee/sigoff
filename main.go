package main

import (
	_ "github.com/kooksee/sigoff/routers"
	"github.com/astaxie/beego"
	"github.com/kooksee/sigoff/conf"
)

func main() {
	cfg := conf.DefaultCfg()
	cfg.InitLogs()
	cfg.InitDb()

	beego.Run()
}

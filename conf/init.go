package conf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"github.com/kooksee/kdb"
)

type cfg struct {
	Debug bool

	db kdb.IKDB
}

func (t *cfg) InitDb() {
	c := kdb.DefaultConfig()
	c.InitKdb()
	t.db = c.GetDb()
}

//日志初始化
func (t *cfg) InitLogs() {
	var level = 7
	if !t.Debug {
		level = 4
	}

	//初始化日志各种配置
	LogsConf := fmt.Sprintf(`{"filename":"kdata/logs/dochub.log","level":%v,"maxlines":5000,"maxsize":0,"daily":true,"maxdays":7}`, level)
	beego.BeeLogger.SetLogger(logs.AdapterFile, LogsConf)
	if t.Debug {
		beego.BeeLogger.SetLogger("console")
	} else {
		//是否异步输出日志
		beego.BeeLogger.Async(1e3)
	}
	beego.BeeLogger.EnableFuncCallDepth(true) //是否显示文件和行号
}

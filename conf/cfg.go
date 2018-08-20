package conf

import (
	"github.com/kooksee/kdb"
	"github.com/astaxie/beego"
)

func (t *cfg) GetDb() kdb.IKDB {
	if t.db == nil {
		panic("please init db")
	}
	return t.db
}

func DefaultCfg() *cfg {
	return &cfg{
		Debug: beego.BConfig.RunMode == "dev",
	}
}

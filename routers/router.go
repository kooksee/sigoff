package routers

import (
	"github.com/kooksee/sigoff/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 提供一个hash写入的接口,能够让系统进行调用
	// 白名单进行过滤，只允许制定的域名才能访问
	// 该hash被发送之后，会显示在用户自己的签名列表中
	// 用户点击确认之后，然后，进行签名操作
	// 签名后的结果发送到服务端
	// 签名是同步的过程，可以等待签名
}

package main

import (
	_ "chsbee/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.SetStaticPath("/static", "static")
	web.Run()
}

package main

import (
	"bit-labs.cn/gin-flex-admin/app/admin"
	_ "bit-labs.cn/gin-flex-admin/docs"
	"bit-labs.cn/owl"
)

func main() {
	var subApps = []owl.SubApp{
		&admin.SubAppAdmin{},
	}
	owl.NewApp(subApps...).Run()
}

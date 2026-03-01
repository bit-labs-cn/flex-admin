package main

import (
	admin "bit-labs.cn/flex-admin/app"
	"bit-labs.cn/owl"
)

func main() {
	var subApps = []owl.SubApp{
		&admin.SubAppAdmin{},
	}
	owl.NewApp(subApps...).WebShell()
}

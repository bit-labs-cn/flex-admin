package main

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
	cms "github.com/guoliang1994/gin-flex-admin/app/cms"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

func main() {

	owl.NewApp("gin-flex-admin").
		BeforeRun(service.MenuStore).
		Run(new(admin.SubAppAdmin), new(cms.SubAppCms))

}

package main

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

func main() {
	owl.Run(new(admin.SubAppAdmin))
}

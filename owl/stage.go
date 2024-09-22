package owl

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl/db"
	"gorm.io/gorm"
)
import "github.com/spf13/cobra"

// SubApp 子应用
type SubApp interface {
	Migrate(db *gorm.DB)
	Seed(db *gorm.DB)
	RegisterRouter(router *gin.Engine)
	RegisterMenu(manager *MenuManager)
	RegisterCommand(command *cobra.Command)
}

var (
	DB        *gorm.DB
	MenuMange *MenuManager
)

func init() {
	DB = db.InitDB(&db.Options{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Driver:   db.Mysql,
		Database: "gfa",
		Schema:   "",
		Charset:  "utf8mb4",
		Query:    "parseTime=True&loc=Local&timeout=3000ms",
	})
}
func Run(apps ...SubApp) {
	router := gin.New()
	cmd := cobra.Command{Use: "owl"}
	MenuMange = &MenuManager{}
	for _, app := range apps {
		app.Migrate(DB)
		app.RegisterRouter(router)
		app.RegisterMenu(MenuMange)
		app.RegisterCommand(&cmd)
		app.Seed(DB)
	}
	router.Run(":8085")
	cmd.Execute()
}

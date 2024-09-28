package owl

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl/db"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)
import "github.com/spf13/cobra"

// SubApp 子应用
type SubApp interface {
	Migrate(db *gorm.DB)
	Seed(db *gorm.DB)
	RegisterRouter()
	RegisterMenu()
	RegisterCommand(command *cobra.Command)
	Name() string
	Construct(application *Application) SubApp
}

type Application struct {
	dig.Container
	db          *gorm.DB
	runDir      string
	baseDir     string
	name        string
	engine      *gin.Engine
	menuManager *MenuManager
	enforcer    casbin.IEnforcer
	beforeRun   []func(a *Application) error
	afterRun    []func(a *Application) error
}

func NewApp(name string) *Application {
	app := &Application{
		name: name,
	}
	app.setPath()
	return app
}

func (i *Application) BeforeRun(fns ...func(a *Application) error) *Application {
	i.beforeRun = fns
	return i
}

// 设置路径
func (i *Application) setPath() {
	var err error
	i.runDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	// 获取当前可执行文件的路径
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	i.baseDir = filepath.Dir(exePath)
}

func (i *Application) initDB() {
	i.db = db.InitDB(&db.Options{
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
func (i *Application) Engine() *gin.Engine {
	return i.engine
}

func (i *Application) Enforcer() casbin.IEnforcer {
	return i.enforcer
}
func (i *Application) MenuManager() *MenuManager {
	return i.menuManager
}
func (i *Application) DB() *gorm.DB {
	return i.db
}

func (i *Application) Routers() []*gin.RouteInfo {
	return i.engine.GetAllRoutes()
}
func (i *Application) Run(apps ...SubApp) {
	i.initDB()

	for _, f := range i.beforeRun {
		err := f(i)
		if err != nil {
			panic(err)
		}
	}

	i.engine = gin.New()
	cmd := cobra.Command{Use: "owl"}
	i.menuManager = &MenuManager{}

	adapter, err := gormadapter.NewAdapterByDB(i.db)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(`
	[request_definition]
	r = sub, act
	
	[policy_definition]
	p = sub, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && r.act == p.act
	`)
	if err != nil {
		return
	}
	i.enforcer, err = casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		app = app.Construct(i) // 构造函数
		app.Migrate(i.db)
		app.RegisterRouter()
		app.RegisterMenu()
		app.RegisterCommand(&cmd)
		app.Seed(i.db)
	}
	if len(i.afterRun) > 0 {
		for _, f := range i.afterRun {
			err = f(i)
			if err != nil {
				panic(err)
			}
		}

	}
	i.engine.Run(":8085")
	cmd.Execute()
}

package admin

import (
	"bit-labs.cn/flex-admin/app/cmd"
	"bit-labs.cn/flex-admin/app/database"
	"bit-labs.cn/flex-admin/app/database/seeder"
	"bit-labs.cn/flex-admin/app/listener"
	"bit-labs.cn/flex-admin/app/provider/jwt"
	"bit-labs.cn/flex-admin/app/route"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract/foundation"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/permission"
	"bit-labs.cn/owl/provider/rabbitmq"
	"bit-labs.cn/owl/provider/redis"
	"bit-labs.cn/owl/provider/router"
	"bit-labs.cn/owl/provider/socketio"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"bit-labs.cn/flex-admin/app/handle/oauth"
	v1 "bit-labs.cn/flex-admin/app/handle/v1"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/flex-admin/app/service"
)

var _ owl.SubApp = (*SubAppAdmin)(nil)

type SubAppAdmin struct {
	app foundation.Application
}

func (i *SubAppAdmin) Name() string {
	return "admin"
}

func (i *SubAppAdmin) Bootstrap() {
	i.app.Invoke(func(db *gorm.DB) {
		database.Migrate(db)
		seeder.InitAllDictData(db)
		listener.Init(i.app)
	})
}

func (i *SubAppAdmin) ServiceProviders() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&permission.GuardProvider{},
		&db.DBServiceProvider{},
		&socketio.SocketIOServiceProvider{},
		&jwt.JwtServiceProvider{},
		&redis.RedisServiceProvider{},
		&rabbitmq.RabbitMQServiceProvider{},
	}
}
func (i *SubAppAdmin) Menu() []*router.Menu {
	return route.InitMenu()
}

func (i *SubAppAdmin) Commands() []*cobra.Command {
	return []*cobra.Command{
		cmd.Version,
	}
}

func (i *SubAppAdmin) RegisterRouters() {
	route.InitApi(i.app, i.Name())
}

func (i *SubAppAdmin) Binds() []any {
	return []any{
		oauth.NewOauthHandle,
		v1.NewApiHandle,
		v1.NewDeptHandle,
		v1.NewDictHandle,
		v1.NewMenuHandle,
		v1.NewRoleHandle,
		v1.NewPositionHandle,
		v1.NewAreaHandle,
		v1.NewLogHandle,
		v1.NewUserHandle,
		service.NewDeptService,
		service.NewDictService,
		service.NewRoleService,
		service.NewLogService,
		service.NewUserService,
		service.NewAreaService,
		repository.NewLogRepository,
		repository.NewDeptRepository,
		repository.NewDictRepository,
		repository.NewRoleRepository,
		repository.NewPositionRepository,
		repository.NewUserRepository,
		repository.NewAreaRepository,
		service.NewPositionService,
	}
}

package admin

import (
	"bit-labs.cn/flex-admin/app/cmd"
	"bit-labs.cn/flex-admin/app/database"
	"bit-labs.cn/flex-admin/app/handle/oauth"
	v1 "bit-labs.cn/flex-admin/app/handle/v1"
	"bit-labs.cn/flex-admin/app/listener"
	admProvider "bit-labs.cn/flex-admin/app/provider"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/flex-admin/app/route"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract/foundation"
	"bit-labs.cn/owl/db"
	"bit-labs.cn/owl/provider"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"net/http"
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
		listener.Init(i.app)
	})
}

func (i *SubAppAdmin) RegisterRouters() {
	route.InitApi(i.app, i.Name())
}

func (i *SubAppAdmin) Binds() []any {
	return []any{
		v1.NewApiHandle,
		v1.NewMenuHandle,

		v1.NewUserHandle,
		service.NewUserService,
		repository.NewUserRepository,

		v1.NewRoleHandle,
		service.NewRoleService,
		repository.NewRoleRepository,

		v1.NewDictHandle,
		service.NewDictService,
		repository.NewDictRepository,

		v1.NewDeptHandle,
		service.NewDeptService,
		repository.NewDeptRepository,

		oauth.NewOauthHandle,

		func() *socketio.Server {
			server := socketio.NewServer(&engineio.Options{
				Transports: []transport.Transport{
					&websocket.Transport{
						CheckOrigin: func(r *http.Request) bool {
							return true
						},
					},
				},
			})
			return server
		},
	}
}
func (i *SubAppAdmin) ServiceProviders() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&provider.GuardProvider{},
		&db.DBServiceProvider{},
		&admProvider.MenuSaveServiceProvider{},
	}
}
func (i *SubAppAdmin) Menu() *owl.Menu {
	return route.InitMenu()
}

func (i *SubAppAdmin) Commands() []*cobra.Command {
	return []*cobra.Command{
		cmd.Version,
		cmd.GenMigrate,
	}
}

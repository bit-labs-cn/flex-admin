package jwt

import (
	_ "embed"

	"bit-labs.cn/owl/contract/foundation"
	"bit-labs.cn/owl/provider/conf"
)

type JwtServiceProvider struct {
	app foundation.Application
}

func (s *JwtServiceProvider) Register() {
	s.app.Register(func(c *conf.Configure) *JWTService {
		var opt JWTOptions
		_ = c.GetConfig("jwt", &opt)

		return NewJWTService(JWTOptions{
			SigningKey: opt.SigningKey,
			Expire:     opt.Expire,
			Issuer:     opt.Issuer,
		})
	})
}
func (s *JwtServiceProvider) Boot() {

}

//go:embed jwt.yaml
var exampleConf string

func (s *JwtServiceProvider) GenerateConf() map[string]string {
	return map[string]string{
		"jwt.yaml": exampleConf,
	}
}

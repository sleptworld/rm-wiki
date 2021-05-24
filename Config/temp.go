package Config

import (
	"github.com/sleptworld/test/DB/plugin"
	"gorm.io/gorm"
)

var (
	AesKey = []byte("ii6Aty+MKr+99Iy5LlT6vEk3Scpn8W7p")
	JWTKey = "fkjljlkasdf"
	Dsn = "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable"
	DBDebug = true
	DBPlugins = []gorm.Plugin{
		&plugin.OpentracingPlugin{},
	}
	ApiVersion = "0.1.0"
)

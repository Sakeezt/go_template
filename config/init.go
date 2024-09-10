package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/integralist/go-findroot/find"
	"github.com/joho/godotenv"
)

type Config struct {
	// Gin config
	GinMode string `env:"GIN_MODE,required"`

	// App config
	AppEnvironment    string `env:"APP_ENV" envDefault:"local"`
	AppName           string `env:"APP_NAME" envDefault:"service"`
	AppTimeZone       string `env:"APP_TIMEZONE" envDefault:"Asia/Bangkok"`
	AppMaxAllowed     int    `env:"APP_MAX_ALLOWED" envDefault:"500"`
	AppLogLevel       int    `env:"APP_LOG_LEVEL" envDefault:"0"`
	AppLogEnableColor bool   `env:"APP_LOG_ENABLE_COLOR" envDefault:"true"`
	AppPort           string `env:"PORT" envDefault:"8080"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT" envDefault:"6831"`
	JaegerLogType   string `env:"JAEGER_LOG_TYPE" envDefault:"stdLogger"`

	// Swagger config
	SwaggerInfoHost     string `env:"SWAGGER_INFO_HOST" envDefault:"localhost:8080"`
	SwaggerInfoBasePath string `env:"SWAGGER_INFO_BASE_PATH" envDefault:"/api/v1"`

	// MongoDB config
	MongoDBEndpoint string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName     string `env:"MONGODB_NAME" envDefault:"service"`
	MongoDBTimeout  int    `env:"MONGODB_TIMEOUT" envDefault:"30"`

	// Redis config
	RedisEndpoint string `env:"REDIS_ENDPOINT" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisTTL      int64  `env:"REDIS_TTL" envDefault:"480"` // minutes
}

func New() *Config {
	conf := &Config{}
	root, _ := find.Repo()
	_ = godotenv.Load(fmt.Sprintf("%s/.env", root.Path))
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	return conf
}

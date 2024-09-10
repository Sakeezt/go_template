package main

import (
	"go-template/app"
	"go-template/config"
	_ "go-template/docs"
	"go-template/service/util/logs"
	"go-template/setup"
)

// @termsOfService  http://swagger.io/terms/
// @contact.name    the developer
// @contact.email   developer@touchtechnologies.co.th
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// load config
	appConfig := config.New()

	// init log
	log := logs.New(appConfig)

	// init setup
	s := setup.New(appConfig, log)

	// setup jaeger
	closer := s.Jaeger()
	defer func() {
		if err := closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	// setup app route and start service
	conf := &app.Config{appConfig, log, s.Logger()}
	app.New(conf).RegisterRoute().Start()
}

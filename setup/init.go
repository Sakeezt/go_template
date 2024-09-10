package setup

import (
	"go-template/config"
	"go-template/service/util/logs"
)

type setup struct {
	appConfig *config.Config
	log       logs.Log
}

func New(appConfig *config.Config, log logs.Log) *setup {
	return &setup{appConfig, log}
}

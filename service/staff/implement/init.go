package implement

import (
	"go-template/service/staff"
	wrp "go-template/service/staff/wrapper"
	"go-template/service/util"
	"go-template/service/util/logs"
	"go-template/service/validator"
)

type implementation struct {
	*StaffServiceConfig
}

type StaffServiceConfig struct {
	Validator    validator.Validator
	RepoStaff    util.Repository
	RepoRedis    util.RepositoryRedis
	UUID         util.UUID
	DateTime     util.DateTime
	FilterString util.FilterString
	Log          logs.Log
}

func New(config *StaffServiceConfig) (service staff.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}

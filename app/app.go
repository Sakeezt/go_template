package app

import (
	"context"

	"go-template/app/staff"
	"go-template/config"
	"go-template/repository/mongodb"
	serviceStaff "go-template/service/staff/implement"
	"go-template/service/util"
	"go-template/service/util/logs"
	serviceValidator "go-template/service/validator"
)

type App struct {
	staff *staff.Controller
}

const (
	collStaff = "staffs"
)

func newApp(appConfig *config.Config, log logs.Log) *App {
	ctx := context.Background()
	uuid, err := util.NewUUID()
	panicIfErr(log, err)
	datetime := util.NewDateTime(appConfig)
	filterString := util.NewFilterString()

	// repositories
	//redisRepo, err := repoRedis.New(ctx, appConfig)
	//panicIfErr(log, err)
	staffRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collStaff))
	panicIfErr(log, err)

	// validators
	validator := serviceValidator.New(&serviceValidator.ValidatorRepository{
		Staff: staffRepo,
	})

	// services
	staffService := serviceStaff.New(&serviceStaff.StaffServiceConfig{
		Validator:    validator,
		RepoStaff:    staffRepo,
		RepoRedis:    nil,
		UUID:         uuid,
		DateTime:     datetime,
		FilterString: filterString,
		Log:          log,
	})

	return &App{
		staff: staff.New(staffService),
	}
}

func mongoDBConfig(appConfig *config.Config, collName string) *mongodb.Config {
	return &mongodb.Config{
		URI:      appConfig.MongoDBEndpoint,
		DBName:   appConfig.MongoDBName,
		CollName: collName,
		Timeout:  appConfig.MongoDBTimeout,
	}
}

func panicIfErr(log logs.Log, err error) {
	if err != nil {
		log.Panic(err)
	}
}

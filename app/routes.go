package app

import (
	"fmt"
	"net/http"
	"os"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginLogRus "github.com/toorop/gin-logrus"

	"go-template/config"
	"go-template/docs"
	"go-template/service/util/logs"
)

type Router struct {
	app       *App
	router    *gin.Engine
	appConfig *config.Config
	log       logs.Log
	logger    *logrus.Logger
}

type Config struct {
	AppConfig *config.Config
	Log       logs.Log
	Logger    *logrus.Logger
}

func New(rc *Config) *Router {
	return &Router{
		app:       newApp(rc.AppConfig, rc.Log),
		router:    gin.New(),
		appConfig: rc.AppConfig,
		log:       rc.Log,
		logger:    rc.Logger,
	}
}

func (r *Router) RegisterRoute() *Router {
	r.routerMiddleware()
	gin.SetMode(r.appConfig.GinMode)

	// api v1
	v1 := r.router.Group(r.appConfig.SwaggerInfoBasePath)
	{
		staff := v1.Group("staffs")
		{
			staff.GET("", r.app.staff.List)
			staff.POST("", r.app.staff.Create)
			staff.GET(":id", r.app.staff.Read)
			staff.PUT(":id", r.app.staff.Update)
			staff.DELETE(":id", r.app.staff.Delete)
		}
	}

	r.routeSwagger()
	r.routeHealthCheck()
	return r
}

func (r *Router) Start() {
	if err := r.router.Run(); err != nil {
		r.log.Panic(err)
	}
}

func (r *Router) routerMiddleware() {
	appENV := os.Getenv("APP_ENV")
	if appENV == "local" {
		r.router.Use(r.ginLogWithConfig())
	} else {
		r.router.Use(ginLogRus.Logger(r.logger))
	}
	r.router.Use(r.corsMiddleware())
	r.router.Use(gin.Recovery())
	r.router.Use(limit.MaxAllowed(r.appConfig.AppMaxAllowed))
}

func (r *Router) routeSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = r.appConfig.AppName
	docs.SwaggerInfo.Description = "API Specification Document"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = r.appConfig.SwaggerInfoHost
	docs.SwaggerInfo.BasePath = r.appConfig.SwaggerInfoBasePath

	docPath := ginSwagger.URL(fmt.Sprintf("//%s/swagger/doc.json", r.appConfig.SwaggerInfoHost))
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docPath))
}

func (r *Router) routeHealthCheck() {
	r.router.GET("/system/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
}

func (r *Router) ginLogWithConfig() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{
			// "/swagger/index.html",
			"/swagger/swagger-ui.css",
			"/swagger/swagger-ui-standalone-preset.js",
			"/swagger/swagger-ui-bundle.js",
			"/swagger/doc.json",
			"/system/health",
			"/favicon.ico",
		},
	})
}

func (r *Router) corsMiddleware() gin.HandlerFunc {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders("Authorization", "User-Agent")
	corsConf.ExposeHeaders = []string{"X-Content-Length"}
	corsConf.AllowHeaders = []string{"*"}
	return cors.New(corsConf)
}

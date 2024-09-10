package setup

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func (s *setup) Logger() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})
	lr.SetLevel(logrus.DebugLevel)
	return lr
}

func (s *setup) Jaeger() io.Closer {
	conf, err := jaegerConf.FromEnv()
	s.panicIfErr(err)

	conf.ServiceName = s.appConfig.AppName
	conf.Sampler.Type = "const"
	conf.Sampler.Param = 1
	conf.Reporter = &jaegerConf.ReporterConfig{LogSpans: true}

	var jLogger jaeger.Logger
	if s.appConfig.JaegerLogType != "nullLogger" {
		jLogger = jaegerLog.StdLogger
	} else {
		jLogger = jaegerLog.NullLogger
	}
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := conf.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
	)
	s.panicIfErr(err)

	opentracing.SetGlobalTracer(tracer)
	return closer
}

func (s *setup) panicIfErr(err error) {
	if err != nil {
		s.log.Panic(err)
	}
}

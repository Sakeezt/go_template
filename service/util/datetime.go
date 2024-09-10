package util

import (
	"go-template/config"

	"github.com/uniplaces/carbon"
)

//go:generate mockery --name=DateTime
type DateTime interface {
	GetNow() *carbon.Carbon
	GetUnixNow() (unix int64)
	ConvertUnixToDateTime(unix int64) (datetime string)
	ConvertUnixToDateTimeWithFormat(unix int64, format string) (datetime string)
}

type dateTime struct {
	appConfig *config.Config
}

func NewDateTime(appConfig *config.Config) (datetime DateTime) {
	return &dateTime{appConfig}
}

func (dt *dateTime) GetNow() (cb *carbon.Carbon) {
	cb, err := carbon.NowInLocation(dt.appConfig.AppTimeZone)
	if err != nil {
		return carbon.Now()
	}
	return cb
}

func (dt *dateTime) GetUnixNow() (unix int64) {
	return dt.GetNow().Unix()
}

func (dt *dateTime) ConvertUnixToDateTime(unix int64) (datetime string) {
	cb, err := carbon.CreateFromTimestamp(unix, dt.appConfig.AppTimeZone)
	if err != nil {
		return ""
	}
	return cb.DateTimeString()
}

func (dt *dateTime) ConvertUnixToDateTimeWithFormat(unix int64, format string) (datetime string) {
	cb, err := carbon.CreateFromTimestamp(unix, dt.appConfig.AppTimeZone)
	if err != nil {
		return ""
	}
	if format != "" {
		return cb.Format(format)
	}
	return cb.DateTimeString()
}

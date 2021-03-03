package global

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DataBaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
	JWTSetting      *setting.JwtSettings
)

package orm

import (
	"errors"
	"github.com/chuangxinyuan/fengchu/component/orm/mysql"
	"github.com/chuangxinyuan/fengchu/component/orm/postgres"
	"github.com/chuangxinyuan/fengchu/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"time"
)

var (
	ErrUnsupportedType = errors.New("unsupported database type")
)

const (
	DriverMysql = "mysql"
	DriverPgSql = "postgres"
)

const (
	Silent string = "silent"
	Error  string = "error"
	Warn   string = "warn"
	Info   string = "info"
)

// Convert convert to gorm logger level
func convert(logLevel string) logger.LogLevel {
	switch logLevel {
	case Silent:
		return logger.Silent
	case Error:
		return logger.Error
	case Warn:
		return logger.Warn
	case Info:
		return logger.Info
	default:
		return logger.Info
	}
}

// DSN database connection configuration
type DSN struct {
	Addr     string
	Database string
	Username string
	Password string
	Options  string
}

// New initialize orm
func New(config *config.Database, logger log.Logger, zLogger *zap.Logger) (db *gorm.DB, cleanup func(), err error) {
	if config == nil {
		return nil, func() {}, nil
	}

	gLogger := zapgorm2.New(zLogger.WithOptions(zap.AddCallerSkip(3)))
	gLogger.SetAsDefault()

	switch config.Driver {
	case DriverMysql:
		db, err = mysql.New(mysql.Config{
			Driver:                    config.Driver,
			Addr:                      config.Address,
			Database:                  config.Database,
			Username:                  config.Username,
			Password:                  config.Password,
			Options:                   config.Options,
			MaxIdleConn:               cast.ToInt(config.MaxIdleConnections),
			MaxOpenConn:               cast.ToInt(config.MaxOpenConnections),
			ConnMaxIdleTime:           config.ConnMaxIdleTime.AsDuration() * time.Second,
			ConnMaxLifeTime:           config.ConnMaxLifeTime.AsDuration() * time.Second,
			Logger:                    gLogger.LogMode(convert(config.LogLevel)),
			Conn:                      nil,
			SkipInitializeWithVersion: false,
			DefaultStringSize:         0,
			DisableDatetimePrecision:  false,
			DontSupportRenameIndex:    false,
			DontSupportRenameColumn:   false,
		})
		if err != nil {
			return
		}
	case DriverPgSql:
		db, err = postgres.New(postgres.Config{
			Driver:               config.Driver,
			Addr:                 config.Address,
			Database:             config.Database,
			Username:             config.Username,
			Password:             config.Password,
			Options:              config.Options,
			MaxIdleConn:          cast.ToInt(config.MaxIdleConnections),
			MaxOpenConn:          cast.ToInt(config.MaxOpenConnections),
			ConnMaxIdleTime:      config.ConnMaxIdleTime.AsDuration() * time.Second,
			ConnMaxLifeTime:      config.ConnMaxLifeTime.AsDuration() * time.Second,
			Logger:               gLogger.LogMode(convert(config.LogLevel)),
			Conn:                 nil,
			PreferSimpleProtocol: false,
		})
		if err != nil {
			return
		}
	default:
		return nil, nil, ErrUnsupportedType
	}

	cleanup = func() {
		log.NewHelper(logger).Info("closing the database resources")

		sqlDB, err := db.DB()
		if err != nil {
			log.NewHelper(logger).Error(err)
		}

		if err := sqlDB.Close(); err != nil {
			log.NewHelper(logger).Error(err)
		}
	}

	return db, cleanup, nil
}

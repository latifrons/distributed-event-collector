package db

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"os"
	"time"
)

type Migrator interface {
	Migrate() error
	SetDB(db *gorm.DB)
}

type DbOperator struct {
	DbType             string
	DSN                string
	DbLog              bool
	MigrateDb          bool
	Db                 *gorm.DB
	MaxIdleConnections int
	MaxOpenConnections int
	LastLogLevel       gormLogger.LogLevel
	Migrator           Migrator
}

func (d *DbOperator) SetLogger(lvl gormLogger.LogLevel) {
	if lvl != d.LastLogLevel {
		d.Db.Config.Logger = d.Db.Config.Logger.LogMode(lvl)
		d.LastLogLevel = lvl
		log.Warn().Int("lvl", int(lvl)).Msg("db log level changed")
	}
}

func (d *DbOperator) InitDefault() (err error) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "01-02 15:04:05.000"}

	defaultLogger := zerolog.New(output).With().Timestamp().Logger()

	if d.DbLog {
		d.LastLogLevel = gormLogger.Info
	} else {
		d.LastLogLevel = gormLogger.Warn
	}

	var db *gorm.DB

	switch d.DbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(d.DSN), &gorm.Config{
			Logger: gormLogger.New(
				&defaultLogger, // IO.writer
				gormLogger.Config{
					SlowThreshold:             time.Second,    // Slow SQL threshold
					LogLevel:                  d.LastLogLevel, // Log minLevel
					IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
					Colorful:                  false,          // Disable color
				},
			),
		})
	//case "postgres":
	//	db, err = gorm.Open(postgres.Open(d.DSN), &gorm.Config{
	//		Logger: gormLogger.New(
	//			&defaultLogger, // IO.writer
	//			gormLogger.Config{
	//				SlowThreshold:             time.Second,    // Slow SQL threshold
	//				LogLevel:                  d.LastLogLevel, // Log minLevel
	//				IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
	//				Colorful:                  false,          // Disable color
	//			},
	//		),
	//	})
	default:
		log.Fatal().Str("dbType", d.DbType).Msg("unsupported db type")
	}

	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return
	}

	d.Db = db

	sqlDB, err := d.Db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(d.MaxIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(d.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if d.MigrateDb {
		d.Migrator.SetDB(d.Db)
		err = d.Migrator.Migrate()
		if err != nil {
			return
		}
	}

	log.Info().Msg("Db inited")
	return
}

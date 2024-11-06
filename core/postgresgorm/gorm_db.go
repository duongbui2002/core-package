package postgresgorm

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
	gromlog "github.com/duongbui2002/core-package/core/logger/external/gorm"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(cfg *GormOptions) (*gorm.DB, error) {
	if cfg.DBName == "" {
		return nil, errors.New("DBName is required in the config.")
	}

	err := createPostgresDB(cfg)
	if err != nil {
		return nil, err
	}

	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
	)

	gormDb, err := gorm.Open(
		gormPostgres.Open(dataSourceName),
		&gorm.Config{
			Logger: gromlog.NewGormCustomLogger(defaultlogger.GetLogger()),
		},
	)
	if err != nil {
		return nil, err
	}

	// add tracing to gorm
	//if cfg.EnableTracing {
	//	err = gormDb.Use(tracing.NewPlugin())
	//}

	return gormDb, nil
}

func createInMemoryDB() (*gorm.DB, error) {
	// https://gorm.io/docs/connecting_to_the_database.html#SQLite
	// https://github.com/glebarez/sqlite
	// https://www.connectionstrings.com/sqlite/
	db, err := gorm.Open(
		sqlite.Open(":memory:"),
		&gorm.Config{
			Logger: gromlog.NewGormCustomLogger(defaultlogger.GetLogger()),
		})

	return db, err
}

func createSQLLiteDB(dbFilePath string) (*gorm.DB, error) {
	// https://gorm.io/docs/connecting_to_the_database.html#SQLite
	// https://github.com/glebarez/sqlite
	// https://www.connectionstrings.com/sqlite/
	gormSQLLiteDB, err := gorm.Open(
		sqlite.Open(dbFilePath),
		&gorm.Config{
			Logger: gromlog.NewGormCustomLogger(defaultlogger.GetLogger()),
		})

	return gormSQLLiteDB, err
}
func NewSQLDB(orm *gorm.DB) (*sql.DB, error) {
	return orm.DB()
}

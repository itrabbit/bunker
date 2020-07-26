package db

import (
	. "github.com/itrabbit/bunker/models"

	"github.com/jinzhu/gorm"

	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//noinspection GoVarAndConstTypeMayBeOmitted
var (
	db       *gorm.DB            = nil
	validate *validator.Validate = validator.New()
)

type Dialect string

const (
	MySQL    Dialect = "mysql"
	MSSQL    Dialect = "mssql"
	SQLite   Dialect = "sqlite3"
	Postgres Dialect = "postgres"
)

func Open(dialect Dialect, args ...interface{}) (err error) {
	db, err = gorm.Open(string(dialect), args...)
	return
}

func AutoMigrate() {
	if db != nil {
		db.AutoMigrate(GetModels()...)
	}
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

package module

import (
	"context"
	"fmt"
	mysql "go.elastic.co/apm/module/apmgormv2/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"interaction-api/config"
)

type DBRepository struct {
	DB *gorm.DB
}

var AppDB *DBRepository

func (r *DBRepository) FindAllWithPagination(db *gorm.DB, ctx context.Context, i interface{}, sql string, offset int, rows int, params ...interface{}) (int64, error) {

	db_ := db.WithContext(ctx)

	type TotalRowData struct {
		Total int64 `gorm:"total"`
	}

	countSql := " /*FORCE_SLAVE*/ " + "SELECT COUNT(1) as total from ( " + sql + " ) count"

	totalRow := TotalRowData{}
	_ = db_.Raw(countSql, params...).Find(&totalRow).Error
	totalRows := totalRow.Total

	sql = " /*FORCE_SLAVE*/ " + sql + "LIMIT ?,?"
	params = append(params, offset, rows)
	err := db_.Raw(sql, params...).Find(i).Error

	return totalRows, err
}

func (r *DBRepository) FindAll(db *gorm.DB, ctx context.Context, i interface{}, sql string, params ...interface{}) error {
	sql = " /*FORCE_SLAVE*/ " + sql
	db_ := db.WithContext(ctx)
	err := db_.Raw(sql, params...).Find(i).Error
	fmt.Println(i)
	return err
}

func (r *DBRepository) FindOne(db *gorm.DB, ctx context.Context, i interface{}, sql string, params ...interface{}) error {
	sql = " /*FORCE_SLAVE*/ " + sql
	db_ := db.WithContext(ctx)
	err := db_.Raw(sql, params...).First(i).Error
	fmt.Println(i)
	return err
}

func init() {
	var logDB logger.Interface
	var dbConnection DBRepository

	// init db log
	if config.AppConfig.DB["db"].DBLogLevel == "ERROR" {
		logDB = logger.Default.LogMode(logger.Error)
	} else {
		logDB = logger.Default.LogMode(logger.Info)
	}

	// build connection string
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", config.AppConfig.DB["db"].User, config.AppConfig.DB["db"].Password, config.AppConfig.DB["db"].Host, config.AppConfig.DB["db"].DB)

	// init db
	if dbConn, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{Logger: logDB, SkipDefaultTransaction: true}); err != nil {
		panic(err)
	} else {
		dbConnection.DB = dbConn
	}

	AppDB = &dbConnection
}

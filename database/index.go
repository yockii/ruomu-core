package database

import (
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"

	logger "github.com/sirupsen/logrus"

	"github.com/yockii/ruomu-core/config"
)

const (
	mysqlConnStringFmt = "%s:%s@tcp(%s:%d)/%s"
	pgConnStringFmt    = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

func initDatabaseDefault() {
	config.DefaultInstance.SetDefault("database.driver", "mysql")
	config.DefaultInstance.SetDefault("database.host", "localhost")
	config.DefaultInstance.SetDefault("database.user", "root")
	config.DefaultInstance.SetDefault("database.password", "root")
	config.DefaultInstance.SetDefault("database.db", "ruomu")
	config.DefaultInstance.SetDefault("database.port", 3306)
	config.DefaultInstance.SetDefault("database.prefix", "t_")
	config.DefaultInstance.SetDefault("database.showSql", false)
}

var DB *xorm.Engine

func Initial() {
	initDatabaseDefault()
	InitDB(
		config.GetString("database.driver"),
		config.GetString("database.host"),
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.db"),
		config.GetInt("database.port"),
		config.GetString("database.prefix"),
		config.GetBool("database.showSql"),
		config.GetString("logger.level"),
	)
}

func InitDB(dbType, host, user, password, dbName string, port int, prefix string, showSql bool, logLevel string) {
	var err error
	DB, err = initDB(dbType, host, user, password, dbName, port)
	if err != nil {
		logger.Fatalf("数据库连接失败! %v", err)
	}
	if err = DB.Ping(); err != nil {
		logger.Fatalf("数据库连接失败! %v", err)
	}
	if prefix != "" {
		DB.SetTableMapper(names.NewPrefixMapper(names.SnakeMapper{}, prefix))
	}

	if showSql {
		DB.ShowSQL(true)
	}
	if logLevel != "" {
		switch strings.ToLower(logLevel) {
		case "error":
			DB.SetLogLevel(log.LOG_ERR)
		case "warn":
			DB.SetLogLevel(log.LOG_WARNING)
		case "info":
			DB.SetLogLevel(log.LOG_INFO)
		case "debug":
			DB.SetLogLevel(log.LOG_DEBUG)
		default:
			DB.SetLogLevel(log.LOG_OFF)
		}
	}

}

func InitDB2(dbDriver, datasource, prefix string, showSql bool, logLevel string) {
	var err error
	DB, err = initDBWithDefine(dbDriver, datasource)
	if err != nil {
		logger.Fatalf("数据库连接失败! %v", err)
	}
	if err = DB.Ping(); err != nil {
		logger.Fatalf("数据库连接失败! %v", err)
	}
	if prefix != "" {
		DB.SetTableMapper(names.NewPrefixMapper(names.SnakeMapper{}, prefix))
	}

	if showSql {
		DB.ShowSQL(true)
	}
	if logLevel != "" {
		switch strings.ToLower(logLevel) {
		case "error":
			DB.SetLogLevel(log.LOG_ERR)
		case "warn":
			DB.SetLogLevel(log.LOG_WARNING)
		case "info":
			DB.SetLogLevel(log.LOG_INFO)
		case "debug":
			DB.SetLogLevel(log.LOG_DEBUG)
		default:
			DB.SetLogLevel(log.LOG_OFF)
		}
	}
}

func initDBWithDefine(driverName, datasourceName string) (*xorm.Engine, error) {
	return xorm.NewEngine(driverName, datasourceName)
}

func initDB(dbType string, host string, user string, password string, dbName string, port int) (*xorm.Engine, error) {
	if dbType == "mysql" {
		return initDBWithDefine("mysql", fmt.Sprintf(
			mysqlConnStringFmt,
			user,
			password,
			host,
			port,
			dbName,
		))
	} else if dbType == "pg" || dbType == "postgres" {
		return initDBWithDefine("postgres", fmt.Sprintf(
			pgConnStringFmt,
			host,
			port,
			user,
			password,
			dbName,
		))
	} else {
		logger.Errorf("数据库初始化失败, 不支持的数据库类型! type=%s, host=%s, user=%s, pwd=%s, db=%s, port=%d", dbType, host, user, password, dbName, port)
		return nil, errors.New("数据库初始化失败, 不支持的数据库类型")
	}
}

func Close() {
	_ = DB.Close()
}

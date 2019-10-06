package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	AppEnv          string = "APP_ENV"
	AppKey                 = "APP_KEY"
	AppName                = "APP_NAME"
	AppHost                = "APP_HOST"
	AppPort                = "APP_PORT"
	AppEndpoint            = "APP_ENDPOINT"
	AppBasepoint           = "APP_BASEPOINT"
	AppRegistryAddr        = "APP_REGISTRY_ADDR"
	AppRegistryPwd         = "APP_REGISTRY_PWD"
	AppTimezone            = "APP_TIMEZONE"
	AppNamespace           = "APP_NAMESPACE"
	AppCluster             = "APP_CLUSTER"

	DBEngine  = "DB_ENGINE"
	DBHost    = "DB_HOST"
	DBPort    = "DB_PORT"
	DBUser    = "DB_USER"
	DBPwd     = "DB_PWD"
	DBName    = "DB_NAME"
	DBSSLMode = "DB_SSL_MODE"
	DBConnStr = "DB_CONN_STR"

	BrokerAddr = "BROKER_ADDR"
	BrokerPwd  = "BROKER_PWD"
)

func (cfg *Config) InitMysql() error {
	driver := Env(DBEngine, "mysql")
	user := Env(DBUser, "root")
	password := Env(DBPwd, "")
	host := Env(DBHost, "127.0.0.1")
	port := Env(DBPort, "3306")
	dbName := Env(DBName, "product")

	dataSource := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"

	db, err := sql.Open(driver, dataSource)

	if err != nil {
		return err
	}

	cfg.DB = db

	return nil
}

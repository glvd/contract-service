package main

import (
	"fmt"
	"net/url"

	"github.com/xormsharp/xorm"
)

const mysqlStatement = "%s:%s@tcp(%s)/%s?loc=%s&charset=%s&parseTime=true"

// MakeInstance ...
func MakeInstance(config DBConfig) (*xorm.Engine, error) {
	if config.DBType == "mysql" {
		return initMysql(config)
	}
	return initSQLite3(config)
}

func liteSource(name string) string {
	return fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL", name)
}

func initSQLite3(config DBConfig) (*xorm.Engine, error) {
	engine, e := xorm.NewEngine(config.DBType, liteSource(config.Schema+".db"))
	if e != nil {
		return nil, e
	}

	return engine, nil
}

func source(config DBConfig) string {
	return fmt.Sprintf(mysqlStatement,
		config.Username,
		config.Password,
		config.Address,
		config.Schema,
		url.QueryEscape("Asia/Shanghai"),
		"utf8mb4")
}

func initMysql(config DBConfig) (*xorm.Engine, error) {
	engine, e := xorm.NewEngine(config.DBType, source(config))
	if e != nil {
		return nil, e
	}

	return engine, nil
}

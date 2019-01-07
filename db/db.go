package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	user     = "root"
	password = "root"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "driversystem"
	charset  = "utf8"
)

var MasterDB *xorm.Engine

func init() {
	if err := initEngine(); err != nil {
		panic(err)
	}

	if err := MasterDB.Ping(); err != nil {
		panic(nil)
	}
}

func fillDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, dbname, charset)
}

func initEngine() error {
	var err error
	MasterDB, err = xorm.NewEngine("mysql", fillDsn())
	if err != nil {
		return err
	}

	return nil
}

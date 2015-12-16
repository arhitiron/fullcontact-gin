package core

import (
	"database/sql"
	"log"

	"github.com/arhitiron/fullcontact-gin/app/models"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var DbMap *gorp.DbMap
var Txn *gorp.Transaction

type Gorp struct {
}

func (c *Gorp) Begin() {
	txn, err := DbMap.Begin()
	if err != nil {
		panic(err)
	}
	Txn = txn
	return
}

func (c *Gorp) Commit() {
	if Txn == nil {
		return
	}

	if err := Txn.Commit(); err != sql.ErrTxDone {
		panic(err)
	}

	Txn = nil
	return
}

func (c *Gorp) Rollback() {
	if Txn == nil {
		return
	}

	if err := Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
	return
}

func InitDb() *gorp.DbMap {
	db, err := sql.Open(Cfg.Database.DriverName, generateDataSourceName())
	checkErr(err, "sql.Open failed")
	DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Cfg.Database.Engine, Cfg.Database.Encoding}}
	DbMap.AddTableWithName(models.Contact{}, "items").SetKeys(true, "Id")
	err = DbMap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return DbMap
}

func generateDataSourceName() string {
	return Cfg.Database.UserName + ":" + Cfg.Database.Password + "@tcp(" + Cfg.Database.Host + ")/" + Cfg.Database.Database + "?charset=" + Cfg.Database.Charset
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

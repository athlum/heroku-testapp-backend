package model

import (
	"os"

	"database/sql"
	"gopkg.in/gorp.v2"
	"time"
)

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (u *User) TableName() string {
	return "user"
}

var db *gorp.DbMap

func DB() *gorp.DbMap {
	return db
}

func Init() error {
	d, err := sql.Open("mysql", os.Getenv("CLEARDB_DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	d.SetConnMaxLifetime(time.Second * 10)
	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(10)

	db = &gorp.DbMap{Db: d, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	u := &User{}
	db.AddTableWithName(u, u.TableName()).SetKeys(true, "id")

	return db.CreateTablesIfNotExists()
}

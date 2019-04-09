package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"sync"
)

var (
	mu      sync.Mutex
	mySqlDb *sql.DB
	Db      sq.StatementBuilderType
)

func Init(dbURL string) (err error) {
	// mySqlDb, err = sql.Open("mysql", "name:password@tcp(ip:port)/dbName?charset=utf8")
	mySqlDb, err = sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}

	err = mySqlDb.Ping()
	if err != nil {
		return err
	}

	mySqlDb.SetMaxOpenConns(10)
	mySqlDb.SetMaxIdleConns(5)

	Db = sq.StatementBuilder.RunWith(mySqlDb)

	return nil
}

func Transaction(call func() error) error {
	mu.Lock()
	tx, err := mySqlDb.Begin()
	if err != nil {
		return err
	}

	Db = Db.RunWith(tx)
	defer func() {
		if txErr := tx.Rollback(); txErr != nil {
			err = txErr
		}
		mu.Unlock()
		Db = Db.RunWith(mySqlDb)
	}()

	err = call()
	if err != nil {
		return err
	}
	if txErr := tx.Commit(); txErr != nil {
		return txErr
	}

	return err
}

func GetSqlDb() *sql.DB{
	return mySqlDb
}

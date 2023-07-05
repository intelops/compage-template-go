package sqls

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sync"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

var o sync.Once

const FileName = "sqlite.db"

type SQLiteClient struct {
	DB *sql.DB
}

var err error
var sqliteClient *SQLiteClient

func InitSqliteDB() (*SQLiteClient, error) {
	o.Do(func() {
		if _, err = os.Stat(FileName); err == nil {
			err = os.Remove(FileName)
		}

		var db *sql.DB
		db, err = sql.Open("sqlite3", FileName)
		sqliteClient = &SQLiteClient{
			DB: db,
		}
	})

	return sqliteClient, err
}

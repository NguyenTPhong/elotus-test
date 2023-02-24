package db

import (
	"database/sql"
	"elotus/config"
	"elotus/database/migration"
	db2 "elotus/package/db"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetTestDB() (*gorm.DB, func()) {
	err := godotenv.Load(basePath + "/../../.env")
	if err != nil {
		panic(err)
	}

	url := config.GetString("TEST_DB_CONN_STR", "")
	if url == "" {
		panic("no test db connection string")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can not connect to database", err.Error())
	}

	databaseName := fmt.Sprintf("test_%v", time.Now().Nanosecond())
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", databaseName))
	if err != nil {
		log.Fatal("can not create test database")
	}

	dbTest, connectErr := db2.NewDatabase(fmt.Sprintf("%v dbname=%v ", url, databaseName), 2, 1, 4)
	if connectErr != nil {
		log.Fatal(errors.New("no db connection"))
	}

	migration.CreateTable(dbTest)

	return dbTest, func() {
		migration.DropTable(dbTest)
		db2.Close(dbTest)
		_, e := db.Exec(fmt.Sprintf("DROP DATABASE %s;", databaseName))
		if e != nil {
			fmt.Println("fail to delete database", e)
		}
		db.Close()
	}
}

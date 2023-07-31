package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func InitDB() error {
	// Ganti "username", "password", "host", "port", dan "dbname" sesuai dengan konfigurasi MySQL Anda
	dataSourceName := "root:@tcp(127.0.0.1:3306)/micropayment"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Ping database untuk memastikan koneksi berhasil
	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}

	DBConn = db
	fmt.Println("Connected to the database")
	return nil
}

func CloseDB() {
	if DBConn != nil {
		DBConn.Close()
	}
}

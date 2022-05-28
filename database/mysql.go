package database

import (
	"database/sql"
	"fmt"
	"simpleService/common"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

type Database struct {
	db *sql.DB
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//DB연결
func Mysql() *Database {

	db, err := sql.Open("mysql", viper.GetString(common.MYSQL_CONNECTION))
	fmt.Println(viper.GetString(common.MYSQL_CONNECTION))
	checkError(err)

	fmt.Println("Database connetion success", db)
	//defer db.Close()
	return &Database{db: db}

}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"simpleService/common"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

var Db *sql.DB

//DB연결
func Mysql() (*sql.DB, error) {
	var err error = nil

	Db, err = sql.Open("mysql", viper.GetString(common.MYSQL_CONNECTION))
	fmt.Println(viper.GetString(common.MYSQL_CONNECTION))
	if err != nil {
		log.Printf("데이터베이스 실패 %s", err.Error())
		return nil, err
	}

	fmt.Println("Database connetion success", Db)

	log.Printf("데이터베이스 연결 성공")
	return Db, nil

}

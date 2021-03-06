package common

import (
	"fmt"

	"github.com/spf13/viper"
)

//셋팅
const (
	MYSQL_CONNECTION = "database.connection_string"
	PORT             = "server.port"
	MONGO_DB_URI     = "database.mongo_connection_string"
	MONGO_DB_NAME    = "database.mongo_db_name"
)

func Setting() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("설정 파일 읽기 실패: %w \n", err))
		return err
	}

	viper.WatchConfig()

	return nil
}

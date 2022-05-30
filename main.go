package main

import (
	"fmt"
	"simpleService/common"
	"simpleService/database"
	"simpleService/routers"
)

func init() {
	//앱셋팅하기 >> 디비 연결준비
	err := common.Setting()

	if err != nil {
		panic(err)
	}

	//디비연결
	database.Mysql()
}
func main() {
	fmt.Println("Im' ready to start!")
	//database.Mysql().AddStore()
	routers.Route()

}

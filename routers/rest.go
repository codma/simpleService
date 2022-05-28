package routers

import (
	"simpleService/database"

	"github.com/gin-gonic/gin"
)

//REST
func Route() {
	//gin 프레임워크를 활용하여 router객체 생성
	router := gin.Default()
	router.GET("/stores", GetStoreList)
	router.Run("localhost:8080")

}

func GetStoreList(c *gin.Context) {
	c.JSON(200, database.Mysql().FindStoreList())
}

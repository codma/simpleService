package routers

import (
	"net/http"
	"simpleService/database"

	"github.com/gin-gonic/gin"
)

//REST
func Route() {
	//gin 프레임워크를 활용하여 router객체 생성
	router := gin.Default()
	router.GET("/stores", GetStoreList)
	router.POST("/stores", AddStore)
	router.PUT("/stores/change_domain", UpdateDomain)
	router.DELETE("/stores", DeleteStore)
	router.Run("localhost:8080")

}

func GetStoreList(c *gin.Context) {
	c.JSON(200, database.Mysql().FindStoreList())
}

func AddStore(c *gin.Context) {
	data, err := database.Mysql().AddStore(c)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, data)
}

func UpdateDomain(c *gin.Context) {
	data, err := database.Mysql().UpdateDomain(c)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func DeleteStore(c *gin.Context) {
	data, err := database.Mysql().DeleteStore(c)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, data)

}

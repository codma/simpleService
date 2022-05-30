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
	c.IndentedJSON(http.StatusCreated, database.Mysql().AddStore(c))
}

func UpdateDomain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Mysql().UpdateDomain(c))
}

func DeleteStore(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Mysql().DeleteStore(c))
}

package routers

import (
	"net/http"
	"simpleService/common"
	"simpleService/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//REST
func Route() {
	//gin 프레임워크를 활용하여 router객체 생성
	router := gin.Default()
	router.GET("/stores", GetStoreList)
	router.POST("/stores", AddStore)
	router.PUT("/stores", UpdateDomain)
	router.DELETE("/stores", DeleteStore)

	port := viper.GetString(common.PORT)
	router.Run("localhost:" + port)

}

func GetStoreList(c *gin.Context) {
	c.JSON(200, database.FindStoreList())
}

func AddStore(c *gin.Context) {
	data, err := database.AddStore(c)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, data)
}

func UpdateDomain(c *gin.Context) {
	data, err := database.UpdateDomain(c)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func DeleteStore(c *gin.Context) {
	data, err := database.DeleteStore(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)

}

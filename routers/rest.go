package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simpleService/common"
	"simpleService/database"
	"simpleService/model"

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

	var storeInfo model.AddOneStore

	data, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(data, &storeInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	result, err := database.AddStore(storeInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
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

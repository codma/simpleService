package main

import (
	"context"
	"fmt"
	"log"
	"simpleService/common"
	"simpleService/database"
	"simpleService/routers"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	//앱셋팅하기 >> 디비 연결준비
	err := common.Setting()
	if err != nil {
		log.Println(err)
	}

	//디비연결
	database.Mysql()
	connectMongo()

}
func main() {
	fmt.Println("Im' ready to start!")

	routers.Route()

}

func connectMongo() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString(common.MONGO_DB_URI)))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database(viper.GetString(common.MONGO_DB_NAME)).Collection("practice")

	return collection
}

package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	database   = "ovms_ali"
	collection = "ovms_bt_product"
)

func main() {
	config := options.Client()
	config.SetHosts([]string{"dds-8vbbb985b9747f641261-pub.mongodb.zhangbei.rds.aliyuncs.com:3717"})
	config.SetAuth(options.Credential{
		Username: "app_user_ali",
		Password: "BmdJ23-1",
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	client, err := mongo.Connect(ctx, config)
	if err != nil {
		fmt.Println("mongo connect error: ", err)
	}
	coll := client.Database(database).Collection(collection)

	// 生成objectId
	objectID, _ := primitive.ObjectIDFromHex("6348f2ad5d0b507fb28fcc95")

	fmt.Println("objectId:", objectID)

	result := coll.FindOne(context.Background(), bson.D{{"_id", objectID}}, options.FindOne())
	if result.Err() != nil {
		fmt.Println("find func error:", result.Err())
	}

	var res map[string]interface{}

	_ = result.Decode(&res)

	for key, value := range res {
		fmt.Printf("res[%v]=%v\n", key, value)
	}

}

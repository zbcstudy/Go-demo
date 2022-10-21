package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
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
	client, err := mongo.NewClient(config)
	if err != nil {
		fmt.Println("mongo connect error: ", err)
	}
	coll := client.Database(database).Collection(collection)

	result := coll.FindOne(context.Background(), bson.D{{"_id", bsoncore.Value{
		Type: bsontype.ObjectID,
		Data: []byte("63455c365d0b503f1022da7b"),
	}}}, options.FindOne())
	if result.Err() != nil {
		fmt.Println("find func error:", result.Err())
	}
	bytes, _ := result.DecodeBytes()
	fmt.Println(string(bytes))
}

package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

const (
	database   = "ovms_ali"
	collection = "ovms_ht_product_sku_eliminated"
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
	// 连接
	client, err := mongo.Connect(ctx, config)
	if err != nil {
		fmt.Println("mongo connect error: ", err)
	}
	watch, err := client.Watch(context.Background(), bson.D{})
	if err != nil {
		return
	}
	watch.ID()
	coll := client.Database(database).Collection(collection)

	// 生成objectId
	objectID, _ := primitive.ObjectIDFromHex("648bdc36626ca006d7567508")
	fmt.Println("objectId:", objectID)

	// 查询单条数据
	result := coll.FindOne(context.Background(), bson.D{{"_id", objectID}}, options.FindOne())
	if result.Err() != nil {
		fmt.Println("find func error:", result.Err())
	}

	// 统计数量
	count, _ := coll.CountDocuments(context.Background(), bson.D{}, options.Count())
	fmt.Println("count:", count)

	// 删除
	//many, _ := coll.DeleteMany(context.Background(), bson.D{{"sellerSku", "AA01-sku-bhtest-6003"}})
	//fmt.Println("delete count:", many)

	// 造数据
	var models []interface{}
	for i := 0; i < 1; i++ {
		em := EliminatedModel{
			supplierId:  "3001",
			sellerId:    "10123",
			code:        "Code-bhtest-6003",
			sku:         "sku-bhtest-6003",
			sellerSku:   "AA01-sku-bhtest-6003",
			platform:    "tmall",
			warehouse:   "TRUNK_30480435",
			operate:     "add",
			brand:       "clarins",
			industry:    "beauty",
			description: "3001供应商",
			alive:       true,
			barCode:     "Barcode-bhtest-6003",
			modifier:    "inventory.control",
			creater:     "inventory.control",
		}
		em.description = em.description + strconv.Itoa(i)
		models = append(models, bson.M{
			"supplierId":  em.supplierId,
			"sellerId":    em.sellerId,
			"code":        em.code,
			"sku":         em.sku,
			"sellerSku":   em.sellerSku,
			"platform":    em.platform,
			"warehouse":   em.warehouse,
			"operate":     em.operate,
			"brand":       em.brand,
			"industry":    em.industry,
			"description": em.description,
			"alive":       em.alive,
			"barCode":     em.barCode,
			"modifier":    em.modifier,
			"creater":     em.creater,
		})
	}
	// 插入
	for count < 1000 {
		insertMany, err := coll.InsertMany(context.Background(), models, options.InsertMany())
		if err != nil {
			fmt.Println("批量插入失败")
		}
		fmt.Println("insertMany:", insertMany)
		count, _ = coll.CountDocuments(context.Background(), bson.D{}, options.Count())
	}

	var res map[string]interface{}

	_ = result.Decode(&res)

	for key, value := range res {
		fmt.Printf("res[%v]=%v\n", key, value)
	}
}

type EliminatedModel struct {
	supplierId  string
	sellerId    string
	code        string
	sku         string
	sellerSku   string
	platform    string
	warehouse   string
	operate     string
	brand       string
	industry    string
	description string
	alive       bool
	barCode     string
	creater     string
	modifier    string
}

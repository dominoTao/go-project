package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "39.106.3.240:6379",
		Password: "",  		// no password set
		DB: 0,				// use default DB
	})
}

//func Get(key string) interface{} {
//	value, err := Client.Get(ctx, key).Result()
//	if err == redis.Nil {
//		fmt.Println(key , " 不存在")
//		return nil
//	}else if err != nil {
//		log.Println(err.Error())
//		return nil
//	}
//	return value
//}
//
//func Set(key, value string) bool {
//	err := Client.Set(ctx, key, value, 0).Err()
//	if err != nil {
//		fmt.Errorf(err.Error())
//		return false
//	}
//	return true
//}
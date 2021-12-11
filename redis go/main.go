package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	ExpireTime = 600
	rdb        *redis.Client
)

func main() {

	rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: ""})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis连接错误")
	}
	a, err := rdb.Exists("lcd").Result()
	if err != nil {
		fmt.Println("判断key存在失败")
		return
	}
	if a == 1 {
		fmt.Println("key 已存在")
	}
	// fmt.Println(time.Duration(ExpireTime) * time.Second)
	//存储key
	err = rdb.Set("lcd", "lcd", time.Duration(ExpireTime)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return
	}

	//获取key
	v, err := rdb.Get("lcd").Result()
	if err != nil {
		fmt.Println("获取key失败")
		return
	}
	fmt.Println(v)

	// 设置过期时间
	err = rdb.Expire("lcd", time.Duration(300)*time.Second).Err()
	if err != nil {
		fmt.Println("设置过期时间失败")
		return
	} else {
		fmt.Println("key 已经失效")
	}

}

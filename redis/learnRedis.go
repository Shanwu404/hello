package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8" // 引入redis包
)

var ctx = context.Background()

func main() {
    // 连接到Redis服务器
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis地址
        Password: "",              // 密码，没有则留空
        DB:       0,               // 使用默认DB
    })

    // 设置键值对
    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    // 获取键的值
    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    // 删除键
    // err = rdb.Del(ctx, "key").Err()
    // if err != nil {
    //     panic(err)
    // }

    // 尝试再次获取已删除的键的值
    // val, err = rdb.Get(ctx, "key").Result()
    // if err == redis.Nil {
    //     fmt.Println("key does not exist")
    // } else if err != nil {
    //     panic(err)
    // } else {
    //     fmt.Println("key", val)
    // }
}

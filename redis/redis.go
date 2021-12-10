package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Dituhui@2017",
		DB:       1,
	})
}

func main() {
	//stringType()
	//hashType()
	//listType()
	mSetType()
}

func mSetType() {
	ctx := context.Background()
	defer rdb.Close()
	err := rdb.HMSet(ctx, "myMSet", "jack", 18, "tom", 20, "jojo", 30).Err()
	if err != nil {
		fmt.Println("save data failed.", err)
		return
	}
	allData := rdb.HGetAll(ctx, "myMSet")
	fmt.Printf("%v", allData.Val())

	err = rdb.HMSet(ctx, "myMSet", "jack", 20).Err()
	if err != nil {
		fmt.Println("save data failed.", err)
		return
	}

	err = rdb.Del(ctx, "myMSet").Err()
	if err != nil {
		fmt.Println("del myMSet failed.", err)
	}
}

func listType() {
	ctx := context.Background()
	defer rdb.Close()
	err := rdb.LPush(ctx, "myList", "a", "b", "c", "c").Err()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 3; i++ {
		pop := rdb.LPop(ctx, "myList")
		fmt.Println("get data:", pop)
	}
	lLen := rdb.LLen(ctx, "myList")
	fmt.Println("list len:", lLen)

	lRange := rdb.LRange(ctx, "myList", 0, -1)
	fmt.Println("get all list data", lRange)

}

func hashType() {
	ctx := context.Background()
	defer rdb.Close()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		res := rdb.HSet(ctx, "myhash", key, fmt.Sprintf("hello%d", i))
		err := res.Err()
		if err != nil {
			fmt.Println("add failed. err", err)
		} else {
			fmt.Printf("data add failed. key:%s\n", key)
		}
	}
	value := rdb.HGet(ctx, "myhash", "key1")
	fmt.Println("get data:", value)

	keys, cursor, err := rdb.HScan(ctx, "myhash", 0, "", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys, cursor)

	res := rdb.HDel(ctx, "myhash", "key1", "key2", "key3")
	result, err := res.Result()
	if err != nil {
		fmt.Println("delete data failed. err:", err)
		return
	}
	fmt.Println(result)

}

func stringType() {
	ctx := context.Background()
	err := rdb.Set(ctx, "tom", 18, 0).Err()
	if err != nil {
		panic(err)
	}
	v := rdb.Get(ctx, "tom")
	fmt.Println(v)

	result, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("Key desc not exits")
	} else if err != nil {
		panic(err)
	}
	rdb.Del(ctx, "tom")
	fmt.Println(result)
}

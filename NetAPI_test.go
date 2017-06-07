package main

import (
	. "clickslash/model"
	. "clickslash/protos"
	. "clickslash/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	//"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestMain(t *testing.T) {

	//testRedis()
	//testData()
	showUserPass()
}

func testData() {
	g_redis := connectRedis()
	_, err := redis.String(g_redis.Do("HGET", "machines", "881"))
	if err != nil {
		fmt.Println("账号不存在")
		fmt.Println(err)
	} else {
		fmt.Println("账号存在")
	}
}

func showUserPass() {
	g_redis := connectRedis()
	str, _ := redis.String(g_redis.Do("HGET", "user1property", "Password"))
	fmt.Println(str)
	fmt.Println(showPassMd5(str))

	str2, _ := redis.String(g_redis.Do("HGET", "user2property", "Password"))
	fmt.Println(str2)
	fmt.Println(showPassMd5(str2))
}

func showPassMd5(password_server string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password_server))
	cipherStr := md5Ctx.Sum(nil)
	password_server = hex.EncodeToString(cipherStr)
	return password_server
}

func testRedis() {
	g_redis := connectRedis()

	id_count, err := redis.String(g_redis.Do("GET", "id_count"))
	if err != nil {
		fmt.Println("get id_count err")
		fmt.Println(err)
		g_redis.Do("SET", "id_count", 0)
		return
	}

	fmt.Println(id_count)

	temp := &TUser{}
	temp.Coin = 5

	//RedisSetStruct(g_redis, "user0property", temp)
	RedisGetStruct(g_redis, "user0property", temp)

	m := make(map[string]interface{})
	Redis2Map(g_redis, "user0property", m)
	fmt.Println(len(m))

}

func connectRedis() redis.Conn {

	fmt.Println("connectRedis")
	var err interface{}
	g_redis, err := redis.Dial("tcp", REDIS_IP)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return g_redis
}

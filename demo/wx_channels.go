package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zsmhub/wx-channels-sdk/apis"
	"log"
	"time"
)

var ctx = context.Background()

var redisDb = redis.NewClient(&redis.Options{
	Addr:     RedisAddr,
	Password: RedisPassword, // no password set
	DB:       RedisDB,       // use default DB
})

// 实现 access_token 的 redis 存取方案
type DcsTokenByRedis struct{}

var _ apis.DcsToken = DcsTokenByRedis{}

func (DcsTokenByRedis) Get(cacheKey string) apis.TokenInfo {
	var tokenInfo apis.TokenInfo
	result, err := redisDb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		_ = json.Unmarshal(result, &tokenInfo)
	} else if err != redis.Nil {
		log.Println("获取 access_token 失败：", err)
	}

	return tokenInfo
}

func (DcsTokenByRedis) Set(cacheKey string, tokenInfo apis.TokenInfo, ttl time.Duration) error {
	data, _ := json.Marshal(tokenInfo)
	err := redisDb.Set(ctx, cacheKey, string(data), ttl).Err()
	if err != nil {
		log.Println("保存 access_token 失败:", err)
	}
	return err
}

func (DcsTokenByRedis) Del(cacheKey string) error {
	return redisDb.Del(ctx, cacheKey).Err()
}

func (DcsTokenByRedis) Lock(cacheKey string, ttl time.Duration) bool {
	if ok, _ := redisDb.SetNX(ctx, cacheKey, 1, ttl).Result(); ok {
		return true
	}
	return false
}

func (DcsTokenByRedis) Unlock(cacheKey string) error {
	return redisDb.Del(ctx, cacheKey).Err()
}

// 日志记录器，可按需改造
type Logger struct{}

var _ apis.Logger = Logger{}

func (Logger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (Logger) Infof(template string, args ...interface{}) {
	fmt.Printf(template, args...)
}

func (Logger) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (Logger) Errorf(template string, args ...interface{}) {
	fmt.Printf(template, args...)
}

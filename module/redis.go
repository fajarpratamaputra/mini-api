package module

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
	"interaction-api/config"
	"time"
)

type RedisWrapper struct {
	Client         *redis.Client
	JwtRedisClient *redis.Client
}

var AppRedis *RedisWrapper

func init() {
	var rdsClient RedisWrapper
	if config.AppConfig.CacheConfig.EnableCache {
		uri := fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port)
		rds := redis.NewClient(&redis.Options{
			Addr:     uri,
			Password: config.AppConfig.Redis.Password,
			DB:       config.AppConfig.Redis.DB,
		})
		rds.AddHook(apmgoredis.NewHook())
		rdsClient.Client = rds
	}
	if config.AppConfig.JwtRedis.Enable {
		uri := fmt.Sprintf("%s:%s", config.AppConfig.JwtRedis.Host, config.AppConfig.JwtRedis.Port)
		rds := redis.NewClient(&redis.Options{
			Addr:     uri,
			Password: config.AppConfig.JwtRedis.Password,
			DB:       config.AppConfig.JwtRedis.DB,
		})
		rds.AddHook(apmgoredis.NewHook())
		rdsClient.JwtRedisClient = rds
	}
	AppRedis = &rdsClient
}

func (r *RedisWrapper) GenerateKey(key string) string {
	return fmt.Sprintf("%s:%s", config.AppConfig.CacheConfig.CachePrefix, key)
}

func (r *RedisWrapper) GetKey(cxt context.Context, key string) (string, error) {
	key = r.GenerateKey(key)
	val, err := r.Client.Get(cxt, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisWrapper) GetParse(ctx context.Context, key string) (map[string]interface{}, error) {
	var valueStr string
	if val, err := r.GetKey(ctx, key); err != nil {
		return nil, err
	} else {
		valueStr = val
	}
	jsonBytes := []byte(valueStr)
	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RedisWrapper) GetKeyCustomClient(cxt context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.Get(cxt, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisWrapper) GetParseCustomClient(ctx context.Context, client *redis.Client, key string) (map[string]interface{}, error) {
	var valueStr string
	if val, err := r.GetKeyCustomClient(ctx, client, key); err != nil {
		return nil, err
	} else {
		valueStr = val
	}
	jsonBytes := []byte(valueStr)
	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RedisWrapper) SetKey(ctx context.Context, value interface{}, key string, duration int) error {
	key = r.GenerateKey(key)
	durationTime := time.Duration(duration) * time.Second
	_, err := r.Client.Set(ctx, key, value, durationTime).Result()
	return err
}

func (r *RedisWrapper) Increment(ctx context.Context, key string) error {
	key = r.GenerateKey(key)
	_, err := r.Client.Incr(ctx, key).Result()
	return err
}

func (r *RedisWrapper) Decrement(ctx context.Context, key string) error {
	key = r.GenerateKey(key)
	_, err := r.Client.Decr(ctx, key).Result()
	return err
}

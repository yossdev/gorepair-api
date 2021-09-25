package cached

// import (
// 	"fmt"
// 	"gorepair-rest-api/config"
// 	"log"
// 	"sync"
// 	"time"

// 	"github.com/go-redis/cache/v8"
// 	"github.com/go-redis/redis/v8"
// )

// ! NOT USING REDIS FOR NOW
// TODO: Uncomment this if using redis and update config/config.go

// var once sync.Once

// type RedisDBInterface interface {
// 	DB() *redis.Client
// 	Cache() *cache.Cache
// 	Close() error
// }

// type redisDB struct {
// 	db    *redis.Client
// 	cache *cache.Cache
// }

// func NewRedisClient() RedisDBInterface {
// 	log.Println("Initialize Redis connection")
// 	host := fmt.Sprintf("%s:%s", config.Get().RedisHost, config.Get().RedisPort)
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     host,
// 		Password: config.Get().RedisPassword, // no password set
// 		DB:       config.Get().RedisDB,       // use default DB
// 	})

// 	ctx := rdb.Context()
// 	ping, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		log.Println("Redis Connection: ", err)
// 	}
// 	log.Println("Redis Connection: ", ping)

// 	cache := cache.New(&cache.Options{
// 		LocalCache: cache.NewTinyLFU(1000, time.Minute),
// 		Redis:      rdb,
// 	})

// 	return &redisDB{
// 		db:    rdb,
// 		cache: cache,
// 	}
// }

// func (c redisDB) DB() *redis.Client {
// 	return c.db
// }

// func (c redisDB) Cache() *cache.Cache {
// 	return c.cache
// }

// func (c redisDB) Close() error {
// 	return c.db.Close()
// }

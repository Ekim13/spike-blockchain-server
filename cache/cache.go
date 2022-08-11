package cache

import (
	"github.com/go-redis/redis"
	logger "github.com/ipfs/go-log"
	"spike-blockchain-server/constants"
)

var log = logger.Logger("cache")

var RedisClient *redis.Client

func Redis() error {
	client := redis.NewClient(
		&redis.Options{
			Addr:       constants.REDIS_ADDR,
			Password:   constants.REDIS_PW,
			MaxRetries: 1,
		})

	_, err := client.Ping().Result()

	if err != nil {
		// log connecting to redis failed
		log.Error("redis init err : ", err)
		panic("redis error")
		return err
	}
	client.SAdd("api_key", "VV6I9K4T1XB9HEZ187XJI51AR2FT8CJ8VV", 0)
	client.SAdd("api_key", "CB3C47F716DC464EF5FB93941FBC8BBD95", 0)
	client.SAdd("api_key", "1ADD650C83877E54E21EB00BF983B85AF9", 0)

	RedisClient = client
	return nil
}

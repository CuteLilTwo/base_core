/*
 * @Description:
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 13:39:06
 * @LastEditTime: 2019-04-30 13:53:31
 */
/*
	建立reids数据库连接
*/
package connect

import (
	"fmt"
	"base_core/config"
	"base_core/log"

	"sync"

	"github.com/go-redis/redis"
)

var (
	RedisEngine *redis.Client
	oneRedis    sync.Once

	DbRedis2 *redis.Client
	DbRedis5 *redis.Client
	DbRedis9 *redis.Client
)

func RedisInit() {
	NewRedis(&config.Conf.Redis)
}

// 单例子模式
func NewRedis(conf *config.RedisConfig) {
	oneRedis.Do(func() {
		log.InfoF("connect to redis %s:%s", conf.Host, conf.Port)
		redisEngine := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port), Password: conf.Pwd, DB: int(conf.Db), PoolSize: 20})
		if redisEngine != nil {
			log.Info("redis had connected")
		} else {
			log.Crit("redis start fail")
		}
		RedisEngine = redisEngine
	},
	)
}

// 连接指定的db库
func NewRedisByDB(dbNum int) *redis.Client {
	c := &config.Conf.Redis
	log.InfoF("connect to redis %s:%s %d", c.Host, c.Port, dbNum)
	redisEngine := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", c.Host, c.Port), Password: c.Pwd, DB: dbNum, PoolSize: 20})
	if redisEngine != nil {
		log.Info("redis had connected")
	} else {
		log.Crit("redis start fail")
	}
	return redisEngine
}

package bootstrap

import (
	"fmt"
	"goex/pkg/config"
	"goex/pkg/redis"
)

// SetupRedis init Redis
func SetupRedis() {
	// build Redis connect
	redis.ConnectRedis(fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}

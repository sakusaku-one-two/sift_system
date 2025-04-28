package redis_client

import (
	// "context"

	env "backend/util"
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context

func GetCLient() *redis.Client {

	if env.GetEnvOrDefault("MODE_IS", "LOCAL_TEST") == "PRODUCTION" {
		//AWSのelastic Cache（redis）を利用
		uri := env.GetEnvOrDefault("REDIS_URI", "")

		opt, err := redis.ParseURL(uri)
		if err != nil {
			panic(err)
		}
		return redis.NewClient(opt)
	}

	//ローカルでの実行環境でのredis(docker)
	address := env.GetEnvOrDefault("REDIS_ADDR", "localhost")
	port := env.GetEnvOrDefault("REDIS_PORT", "6379")
	ADDR := address + ":" + port
	password := env.GetEnvOrDefault("REDIS_PASSWORD", "my_password")
	pool_size := env.GetEnvOrDefaultToInt("REDIS_POOL_SIZE", "100")
	user_name := env.GetEnvOrDefault("REDIS_USER", "backendUser")

	client := redis.NewClient(&redis.Options{
		Addr:     ADDR,
		Password: password, // No password set
		Username: user_name,
		DB:       0, // Use default DB
		PoolSize: pool_size,
	})

	ctx = context.Background()

	pong, err := client.Ping(ctx).Result()

	if err != nil {
		env.Print("接続が失敗しました。", err, pong)
	}

	return client
}

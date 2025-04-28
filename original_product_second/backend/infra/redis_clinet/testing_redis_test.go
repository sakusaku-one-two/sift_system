package redis_client

import (
	"backend/util"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {

	const key = "test-key"
	rdb := GetCLient()

	defer rdb.Close()

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	go monitorConnections(ctx, rdb)

	err = rdb.Set(ctx, key, "test-value", 5*time.Minute).Err()

	if err != nil {
		util.Print("失敗しました", err)
	}

	//キーを取得
	val, err := rdb.Get(ctx, key).Result()

	util.Print(val, err)

	time.Sleep(time.Second * 20)
}

func monitorConnections(ctx context.Context, rdb *redis.Client) {
	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	for {
		<-ticker.C
		stats := rdb.PoolStats()
		log.Printf("Redis接続状態:総数=%d,アイドル：%d",
			stats.TotalConns, stats.IdleConns,
		)

		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}

func Test_setValue(t *testing.T) {
	rdb_service, err := NewRedisService()
	if err != nil {
		util.Print(err)
		return
	}

	util.Print(rdb_service.RDB)

}

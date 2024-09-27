package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	Address    string
	Password   string
	Db         int
	RootFolder string
	Rdb        *redis.Client
}

func (r *RedisCache) Init() {
	r.Rdb = redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.Db,
	})
}

// GetExclusiveLock will try to get a redis lock and release it after some time
// it is necessary to get a lock and lock it for a while to globally schedule a cron.
func (r *RedisCache) GetExclusiveLock(ctx context.Context, group string, name string, lockTime time.Duration) (success bool, err error) {
	nx := r.Rdb.SetNX(ctx, r.WrapKey(group+"-"+name), 1, lockTime)
	err = nx.Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return
		}
		err = nil
		success = false
		return
	}
	return nx.Val(), nil
}

func (r *RedisCache) ReleaseExclusiveLock(ctx context.Context, group string, name string) (err error) {
	nx := r.Rdb.Del(ctx, r.WrapKey(group+"-"+name))
	err = nx.Err()
	return
}

func (r *RedisCache) WrapKey(key string) string {
	return r.RootFolder + ":" + key
}

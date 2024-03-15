package other

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis"
)

func LockFeature(
	ctx context.Context,
	redisClient *redis.Client,
	duration time.Duration,
	keyRedis string,
	retryTimes int) (
	err error,
) {
	if redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	if retryTimes <= 0 {
		retryTimes = 1
	}

	res := redsync.New(goredis.NewPool(redisClient))
	mutex := res.NewMutex(
		keyRedis,
		redsync.WithExpiry(duration),
		redsync.WithTries(retryTimes),
		redsync.WithRetryDelay(time.Second), // TODO: move to option, now just quick update to don't need to change all of current implement
	)

	if err = mutex.LockContext(ctx); err != nil {
		if errors.Is(err, redsync.ErrFailed) {
			return errors.New("cannot be empty")
		}

		return err
	}

	return nil
}

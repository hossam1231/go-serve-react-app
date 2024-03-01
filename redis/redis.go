//go:generate go run github.com/99designs/gqlgen generate
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hossam1231/logger-go-pkg"
)

// Connect connects to a Redis instance and returns a Redis client.
func Connect(ctx context.Context) (*redis.Client, error) {

	// Connect to Redis
	opt, err := redis.ParseURL("redis://default:4d46a54bd854423f8b7be606c255f11f@eu1-lenient-heron-38709.upstash.io:38709")
	if err != nil {
		logger.Error("failed to parse Redis URL: %v", err)
		return nil, err
	}
	client := redis.NewClient(opt)

	logger.Success("Redis client connected")

	return client, nil
}

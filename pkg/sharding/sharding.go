package sharding

import (
	"gorm.io/sharding"
)

func NewSharding(shardingKey string, shardingNumber uint, tableName string) *sharding.Sharding {
	handler := sharding.Register(sharding.Config{
		ShardingKey:         shardingKey,
		NumberOfShards:      shardingNumber,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, tableName)
	return handler
}

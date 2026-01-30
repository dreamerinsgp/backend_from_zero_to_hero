package cache

import (
	"cache-demo/model"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// BloomFilterKey 布隆过滤器在Redis中的key
	BloomFilterKey = "user_bloom_filter"
	// BloomFilterBits 布隆过滤器位数组大小（100万，误判率约1%）
	BloomFilterBits = 1000000
)

// UserCacheWithBloom 支持布隆过滤器的用户缓存接口
type UserCacheWithBloom interface {
	GetUser(id int64) (*model.User, error)
	SetUser(user *model.User, expireSeconds int) error
	DeleteUser(id int64) error
	AddToBloomFilter(id int64) error
	ExistsInBloomFilter(id int64) (bool, error)
}

// userCacheWithBloom 支持布隆过滤器的用户缓存实现
type userCacheWithBloom struct {
	rds         *redis.Redis
	bloomFilter *bloom.Filter
}

// NewUserCacheWithBloom 创建支持布隆过滤器的用户缓存实例
func NewUserCacheWithBloom(rds *redis.Redis) UserCacheWithBloom {
	bloomFilter := bloom.New(rds, BloomFilterKey, BloomFilterBits)
	return &userCacheWithBloom{
		rds:         rds,
		bloomFilter: bloomFilter,
	}
}

// getUserKey 生成用户缓存Key
func (c *userCacheWithBloom) getUserKey(id int64) string {
	return fmt.Sprintf("%s%d", UserCacheKeyPrefix, id)
}

// GetUser 从Redis获取用户信息
func (c *userCacheWithBloom) GetUser(id int64) (*model.User, error) {
	key := c.getUserKey(id)

	// 从Redis获取数据
	val, err := c.rds.Get(key)
	if err != nil {
		return nil, err
	}

	// 反序列化JSON数据
	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, fmt.Errorf("反序列化用户数据失败: %w", err)
	}

	return &user, nil
}

// SetUser 设置用户信息到Redis
func (c *userCacheWithBloom) SetUser(user *model.User, expireSeconds int) error {
	if user == nil {
		return fmt.Errorf("用户数据不能为空")
	}

	key := c.getUserKey(user.ID)

	// 序列化为JSON
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("序列化用户数据失败: %w", err)
	}

	// 设置到Redis，带过期时间
	if expireSeconds <= 0 {
		expireSeconds = DefaultExpireSeconds
	}

	err = c.rds.Setex(key, string(data), expireSeconds)
	if err != nil {
		return fmt.Errorf("设置缓存失败: %w", err)
	}

	return nil
}

// DeleteUser 删除用户缓存
func (c *userCacheWithBloom) DeleteUser(id int64) error {
	key := c.getUserKey(id)
	_, err := c.rds.Del(key)
	return err
}

// AddToBloomFilter 添加用户ID到布隆过滤器
func (c *userCacheWithBloom) AddToBloomFilter(id int64) error {
	key := fmt.Sprintf("user:%d", id)
	return c.bloomFilter.Add([]byte(key))
}

// ExistsInBloomFilter 检查用户ID是否在布隆过滤器中
func (c *userCacheWithBloom) ExistsInBloomFilter(id int64) (bool, error) {
	key := fmt.Sprintf("user:%d", id)
	return c.bloomFilter.Exists([]byte(key))
}

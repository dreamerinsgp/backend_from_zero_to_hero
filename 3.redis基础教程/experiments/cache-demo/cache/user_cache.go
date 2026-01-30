package cache

import (
	"cache-demo/model"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// UserCacheKeyPrefix 用户缓存Key前缀
	UserCacheKeyPrefix = "user:"
	// DefaultExpireSeconds 默认过期时间（5分钟）
	DefaultExpireSeconds = 300
)

// UserCache 用户缓存接口
type UserCache interface {
	GetUser(id int64) (*model.User, error)
	SetUser(user *model.User, expireSeconds int) error
	DeleteUser(id int64) error
}

// userCache 用户缓存实现
type userCache struct {
	rds *redis.Redis
}

// NewUserCache 创建用户缓存实例
func NewUserCache(rds *redis.Redis) UserCache {
	return &userCache{rds: rds}
}

// getUserKey 生成用户缓存Key
func getUserKey(id int64) string {
	return fmt.Sprintf("%s%d", UserCacheKeyPrefix, id)
}

// GetUser 从Redis获取用户信息
func (c *userCache) GetUser(id int64) (*model.User, error) {
	key := getUserKey(id)

	// 从Redis获取数据
	val, err := c.rds.Get(key)
	if err != nil {
		// Redis中没有数据或连接错误
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
func (c *userCache) SetUser(user *model.User, expireSeconds int) error {
	if user == nil {
		return fmt.Errorf("用户数据不能为空")
	}

	key := getUserKey(user.ID)

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
func (c *userCache) DeleteUser(id int64) error {
	key := getUserKey(id)
	_, err := c.rds.Del(key)
	return err
}

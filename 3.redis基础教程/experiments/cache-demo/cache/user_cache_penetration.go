package cache

import (
	"cache-demo/model"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// NullCacheValue 空值缓存的标记值
	NullCacheValue = "NULL"
	// NullCacheExpireSeconds 空值缓存过期时间（60秒）
	NullCacheExpireSeconds = 60
)

// UserCacheWithPenetration 支持缓存穿透防护的用户缓存接口
type UserCacheWithPenetration interface {
	GetUser(id int64) (*model.User, error)
	SetUser(user *model.User, expireSeconds int) error
	SetNullUser(id int64) error
	IsNullCache(id int64) (bool, error)
	DeleteUser(id int64) error
}

// userCacheWithPenetration 支持缓存穿透防护的用户缓存实现
type userCacheWithPenetration struct {
	rds *redis.Redis
}

// NewUserCacheWithPenetration 创建支持缓存穿透防护的用户缓存实例
func NewUserCacheWithPenetration(rds *redis.Redis) UserCacheWithPenetration {
	return &userCacheWithPenetration{rds: rds}
}

// getUserKey 生成用户缓存Key（复用原有函数）
func (c *userCacheWithPenetration) getUserKey(id int64) string {
	return fmt.Sprintf("%s%d", UserCacheKeyPrefix, id)
}

// GetUser 从Redis获取用户信息（支持空值缓存）
func (c *userCacheWithPenetration) GetUser(id int64) (*model.User, error) {
	key := c.getUserKey(id)

	// 从Redis获取数据
	val, err := c.rds.Get(key)
	if err != nil {
		// Redis中没有数据或连接错误
		return nil, err
	}

	// 检查是否是空值缓存
	if val == NullCacheValue {
		// 返回特殊错误，表示是空值缓存
		return nil, ErrNullCache
	}

	// 反序列化JSON数据
	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, fmt.Errorf("反序列化用户数据失败: %w", err)
	}

	return &user, nil
}

// SetUser 设置用户信息到Redis
func (c *userCacheWithPenetration) SetUser(user *model.User, expireSeconds int) error {
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

// SetNullUser 设置空值缓存（用于防止缓存穿透）
func (c *userCacheWithPenetration) SetNullUser(id int64) error {
	key := c.getUserKey(id)
	err := c.rds.Setex(key, NullCacheValue, NullCacheExpireSeconds)
	if err != nil {
		return fmt.Errorf("设置空值缓存失败: %w", err)
	}
	return nil
}

// IsNullCache 检查是否是空值缓存
func (c *userCacheWithPenetration) IsNullCache(id int64) (bool, error) {
	key := c.getUserKey(id)
	val, err := c.rds.Get(key)
	if err != nil {
		return false, err
	}
	return val == NullCacheValue, nil
}

// DeleteUser 删除用户缓存（包括空值缓存）
func (c *userCacheWithPenetration) DeleteUser(id int64) error {
	key := c.getUserKey(id)
	_, err := c.rds.Del(key)
	return err
}

// ErrNullCache 空值缓存错误（用于标识空值缓存命中）
var ErrNullCache = fmt.Errorf("空值缓存命中")

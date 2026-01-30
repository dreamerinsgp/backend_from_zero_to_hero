package cache

import (
	"cache-demo/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// AvalancheBaseExpireSeconds 缓存雪崩测试的基础过期时间（60秒，便于测试观察）
	AvalancheBaseExpireSeconds = 60
	// AvalancheRandomRangePercent 随机范围百分比（10%）
	AvalancheRandomRangePercent = 10
)

// UserCacheWithAvalanche 支持缓存雪崩优化的用户缓存接口
type UserCacheWithAvalanche interface {
	GetUser(id int64) (*model.User, error)
	SetUserWithRandomExpire(user *model.User, baseExpireSeconds int) error
	SetUserWithFixedExpire(user *model.User, expireSeconds int) error
	DeleteUser(id int64) error
}

// userCacheWithAvalanche 支持缓存雪崩优化的用户缓存实现
type userCacheWithAvalanche struct {
	rds *redis.Redis
}

// NewUserCacheWithAvalanche 创建支持缓存雪崩优化的用户缓存实例
func NewUserCacheWithAvalanche(rds *redis.Redis) UserCacheWithAvalanche {
	return &userCacheWithAvalanche{rds: rds}
}

// getUserKey 生成用户缓存Key
func (c *userCacheWithAvalanche) getUserKey(id int64) string {
	return fmt.Sprintf("%s%d", UserCacheKeyPrefix, id)
}

// GetUser 从Redis获取用户信息
func (c *userCacheWithAvalanche) GetUser(id int64) (*model.User, error) {
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

// SetUserWithRandomExpire 设置用户信息到Redis（使用随机过期时间，防止缓存雪崩）
func (c *userCacheWithAvalanche) SetUserWithRandomExpire(user *model.User, baseExpireSeconds int) error {
	if user == nil {
		return fmt.Errorf("用户数据不能为空")
	}

	key := c.getUserKey(user.ID)

	// 序列化为JSON
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("序列化用户数据失败: %w", err)
	}

	// 计算随机过期时间：基础过期时间 + 随机值（0-10%范围）
	if baseExpireSeconds <= 0 {
		baseExpireSeconds = DefaultExpireSeconds
	}

	// 随机范围：基础过期时间的0-10%
	randomRange := baseExpireSeconds * AvalancheRandomRangePercent / 100
	if randomRange == 0 {
		randomRange = 1 // 至少1秒的随机范围
	}
	randomExpire := rand.Intn(randomRange)
	expire := baseExpireSeconds + randomExpire

	err = c.rds.Setex(key, string(data), expire)
	if err != nil {
		return fmt.Errorf("设置缓存失败: %w", err)
	}

	return nil
}

// SetUserWithFixedExpire 设置用户信息到Redis（使用固定过期时间，用于模拟缓存雪崩）
func (c *userCacheWithAvalanche) SetUserWithFixedExpire(user *model.User, expireSeconds int) error {
	if user == nil {
		return fmt.Errorf("用户数据不能为空")
	}

	key := c.getUserKey(user.ID)

	// 序列化为JSON
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("序列化用户数据失败: %w", err)
	}

	// 设置到Redis，带固定过期时间
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
func (c *userCacheWithAvalanche) DeleteUser(id int64) error {
	key := c.getUserKey(id)
	_, err := c.rds.Del(key)
	return err
}

// GetRandomExpireTime 计算随机过期时间（用于测试和日志）
func GetRandomExpireTime(baseExpireSeconds int) int {
	if baseExpireSeconds <= 0 {
		baseExpireSeconds = DefaultExpireSeconds
	}
	randomRange := baseExpireSeconds * AvalancheRandomRangePercent / 100
	if randomRange == 0 {
		randomRange = 1
	}
	randomExpire := rand.Intn(randomRange)
	return baseExpireSeconds + randomExpire
}

// 初始化随机数种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

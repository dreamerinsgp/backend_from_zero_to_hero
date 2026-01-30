package service

import (
	"cache-demo/cache"
	"cache-demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// ExpireMode 过期时间模式
type ExpireMode int

const (
	// FixedExpire 固定过期时间（模拟缓存雪崩）
	FixedExpire ExpireMode = iota
	// RandomExpire 随机过期时间（解决缓存雪崩）
	RandomExpire
)

// userServiceWithAvalanche 支持缓存雪崩优化的用户服务实现
type userServiceWithAvalanche struct {
	repo       model.UserRepo
	cache      cache.UserCacheWithAvalanche
	expireMode ExpireMode
	baseExpire int
}

// NewUserServiceWithAvalanche 创建支持缓存雪崩优化的用户服务实例
func NewUserServiceWithAvalanche(repo model.UserRepo, cache cache.UserCacheWithAvalanche, mode ExpireMode, baseExpire int) UserService {
	return &userServiceWithAvalanche{
		repo:       repo,
		cache:      cache,
		expireMode: mode,
		baseExpire: baseExpire,
	}
}

// GetUserByID 根据ID获取用户（支持固定/随机过期时间）
func (s *userServiceWithAvalanche) GetUserByID(id int64) (*model.User, error) {
	// 1. 先查缓存
	user, err := s.cache.GetUser(id)
	if err == nil && user != nil {
		log.Printf("[缓存命中] user_id=%d, username=%s", id, user.Username)
		return user, nil
	}

	// 2. 缓存未命中，查数据库
	log.Printf("[缓存未命中] user_id=%d, 查询数据库", id)
	user, err = s.repo.FindByID(id)
	if err != nil {
		// 检查是否是记录不存在的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[用户不存在] user_id=%d", id)
			return nil, fmt.Errorf("用户不存在: user_id=%d", id)
		}
		// 其他数据库错误
		log.Printf("[数据库查询失败] user_id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 3. 写入缓存（根据模式选择固定或随机过期时间）
	if user == nil {
		log.Printf("[用户不存在] user_id=%d", id)
		return nil, fmt.Errorf("用户不存在: user_id=%d", id)
	}

	switch s.expireMode {
	case FixedExpire:
		// 固定过期时间（模拟缓存雪崩）
		if err := s.cache.SetUserWithFixedExpire(user, s.baseExpire); err != nil {
			log.Printf("[缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
		} else {
			log.Printf("[缓存写入成功] user_id=%d, username=%s, expire=%d秒 (固定)", id, user.Username, s.baseExpire)
		}

	case RandomExpire:
		// 随机过期时间（解决缓存雪崩）
		// 注意：实际过期时间在SetUserWithRandomExpire内部计算，这里只是估算用于日志
		if err := s.cache.SetUserWithRandomExpire(user, s.baseExpire); err != nil {
			log.Printf("[缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
		} else {
			// 估算过期时间范围用于日志（实际值在缓存层计算）
			maxExpire := s.baseExpire + s.baseExpire*cache.AvalancheRandomRangePercent/100
			log.Printf("[缓存写入成功] user_id=%d, username=%s, expire=%d-%d秒 (随机)", id, user.Username, s.baseExpire, maxExpire)
		}
	}

	return user, nil
}

// CreateUser 创建用户
func (s *userServiceWithAvalanche) CreateUser(user *model.User) error {
	log.Printf("[创建用户] username=%s", user.Username)

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	log.Printf("[创建用户成功] user_id=%d, username=%s", user.ID, user.Username)
	return nil
}

// UpdateUser 更新用户
func (s *userServiceWithAvalanche) UpdateUser(user *model.User) error {
	log.Printf("[更新用户] user_id=%d", user.ID)

	// 1. 更新数据库
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	// 2. 更新缓存（根据模式选择固定或随机过期时间）
	switch s.expireMode {
	case FixedExpire:
		if err := s.cache.SetUserWithFixedExpire(user, s.baseExpire); err != nil {
			log.Printf("[缓存更新失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
		} else {
			log.Printf("[缓存更新成功] user_id=%d (固定过期时间)", user.ID)
		}

	case RandomExpire:
		if err := s.cache.SetUserWithRandomExpire(user, s.baseExpire); err != nil {
			log.Printf("[缓存更新失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
		} else {
			log.Printf("[缓存更新成功] user_id=%d (随机过期时间)", user.ID)
		}
	}

	return nil
}

// DeleteUser 删除用户
func (s *userServiceWithAvalanche) DeleteUser(id int64) error {
	log.Printf("[删除用户] user_id=%d", id)

	// 1. 删除数据库记录
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	// 2. 删除缓存
	if err := s.cache.DeleteUser(id); err != nil {
		log.Printf("[缓存删除失败] user_id=%d, error=%v (不影响业务逻辑)", id, err)
	} else {
		log.Printf("[缓存删除成功] user_id=%d", id)
	}

	return nil
}

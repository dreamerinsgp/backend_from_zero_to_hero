package service

import (
	"cache-demo/cache"
	"cache-demo/model"
	"fmt"
	"log"
)

// UpdateStrategy 缓存更新策略类型
type UpdateStrategy int

const (
	// UpdateCache 更新缓存策略（读多写少）
	UpdateCache UpdateStrategy = iota
	// DeleteCache 删除缓存策略（写多读少、数据一致性要求高）
	DeleteCache
)

// userServiceWithStrategy 带策略的用户服务实现
type userServiceWithStrategy struct {
	repo     model.UserRepo
	cache    cache.UserCache
	strategy UpdateStrategy
}

// NewUserServiceWithStrategy 创建带策略的用户服务实例
func NewUserServiceWithStrategy(repo model.UserRepo, cache cache.UserCache, strategy UpdateStrategy) UserService {
	return &userServiceWithStrategy{
		repo:     repo,
		cache:    cache,
		strategy: strategy,
	}
}

// GetUserByID 根据ID获取用户（Cache-Aside 模式）
func (s *userServiceWithStrategy) GetUserByID(id int64) (*model.User, error) {
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
		log.Printf("[数据库查询失败] user_id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 如果用户不存在，返回nil
	if user == nil {
		log.Printf("[用户不存在] user_id=%d", id)
		return nil, fmt.Errorf("用户不存在: user_id=%d", id)
	}

	// 3. 写入缓存（设置过期时间 5 分钟）
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
	} else {
		log.Printf("[缓存写入成功] user_id=%d, username=%s, expire=%d秒", id, user.Username, cache.DefaultExpireSeconds)
	}

	return user, nil
}

// CreateUser 创建用户
func (s *userServiceWithStrategy) CreateUser(user *model.User) error {
	log.Printf("[创建用户] username=%s", user.Username)

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	log.Printf("[创建用户成功] user_id=%d, username=%s", user.ID, user.Username)
	return nil
}

// UpdateUser 更新用户（根据策略选择更新缓存或删除缓存）
func (s *userServiceWithStrategy) UpdateUser(user *model.User) error {
	log.Printf("[更新用户] user_id=%d, 策略=%v", user.ID, s.strategy)

	// 1. 更新数据库
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	// 2. 根据策略更新或删除缓存
	switch s.strategy {
	case UpdateCache:
		// 策略1：更新缓存（读多写少）
		if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
			log.Printf("[缓存更新失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
		} else {
			log.Printf("[缓存更新成功] user_id=%d (策略: 更新缓存)", user.ID)
		}

	case DeleteCache:
		// 策略2：删除缓存（写多读少、数据一致性要求高）
		if err := s.cache.DeleteUser(user.ID); err != nil {
			log.Printf("[缓存删除失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
		} else {
			log.Printf("[缓存删除成功] user_id=%d (策略: 删除缓存)", user.ID)
		}
	}

	return nil
}

// DeleteUser 删除用户
func (s *userServiceWithStrategy) DeleteUser(id int64) error {
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

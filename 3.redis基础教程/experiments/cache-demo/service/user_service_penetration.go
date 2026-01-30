package service

import (
	"cache-demo/cache"
	"cache-demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// userServiceWithPenetration 支持缓存穿透防护的用户服务实现
type userServiceWithPenetration struct {
	repo  model.UserRepo
	cache cache.UserCacheWithPenetration
}

// NewUserServiceWithPenetration 创建支持缓存穿透防护的用户服务实例
func NewUserServiceWithPenetration(repo model.UserRepo, cache cache.UserCacheWithPenetration) UserService {
	return &userServiceWithPenetration{
		repo:  repo,
		cache: cache,
	}
}

// GetUserByID 根据ID获取用户（支持空值缓存，防止缓存穿透）
func (s *userServiceWithPenetration) GetUserByID(id int64) (*model.User, error) {
	// 1. 先查缓存
	user, err := s.cache.GetUser(id)
	if err == nil && user != nil {
		log.Printf("[缓存命中] user_id=%d, username=%s", id, user.Username)
		return user, nil
	}

	// 2. 检查是否是空值缓存
	if errors.Is(err, cache.ErrNullCache) {
		log.Printf("[空值缓存命中] user_id=%d (防止缓存穿透)", id)
		return nil, fmt.Errorf("用户不存在: user_id=%d", id)
	}

	// 3. 缓存未命中，查数据库
	log.Printf("[缓存未命中] user_id=%d, 查询数据库", id)
	user, err = s.repo.FindByID(id)
	if err != nil {
		// 检查是否是记录不存在的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户不存在，设置空值缓存（防止缓存穿透）
			log.Printf("[用户不存在] user_id=%d, 设置空值缓存 (expire=%d秒)", id, cache.NullCacheExpireSeconds)
			if err := s.cache.SetNullUser(id); err != nil {
				log.Printf("[空值缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
			} else {
				log.Printf("[空值缓存写入成功] user_id=%d, expire=%d秒", id, cache.NullCacheExpireSeconds)
			}
			return nil, fmt.Errorf("用户不存在: user_id=%d", id)
		}
		// 其他数据库错误（如连接失败等）
		log.Printf("[数据库查询失败] user_id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 4. 处理查询结果（正常情况下 user 不为 nil，但为了安全起见还是检查一下）
	if user == nil {
		// 用户不存在，设置空值缓存（防止缓存穿透）
		log.Printf("[用户不存在] user_id=%d, 设置空值缓存 (expire=%d秒)", id, cache.NullCacheExpireSeconds)
		if err := s.cache.SetNullUser(id); err != nil {
			log.Printf("[空值缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
		} else {
			log.Printf("[空值缓存写入成功] user_id=%d, expire=%d秒", id, cache.NullCacheExpireSeconds)
		}
		return nil, fmt.Errorf("用户不存在: user_id=%d", id)
	}

	// 5. 用户存在，写入正常缓存
	log.Printf("[用户存在] user_id=%d, username=%s, 写入缓存", id, user.Username)
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
	} else {
		log.Printf("[缓存写入成功] user_id=%d, username=%s, expire=%d秒", id, user.Username, cache.DefaultExpireSeconds)
	}

	return user, nil
}

// CreateUser 创建用户
func (s *userServiceWithPenetration) CreateUser(user *model.User) error {
	log.Printf("[创建用户] username=%s", user.Username)

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	// 如果之前有空值缓存，需要删除（因为现在用户已存在）
	if err := s.cache.DeleteUser(user.ID); err != nil {
		log.Printf("[删除空值缓存失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
	}

	log.Printf("[创建用户成功] user_id=%d, username=%s", user.ID, user.Username)
	return nil
}

// UpdateUser 更新用户
func (s *userServiceWithPenetration) UpdateUser(user *model.User) error {
	log.Printf("[更新用户] user_id=%d", user.ID)

	// 1. 更新数据库
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	// 2. 更新缓存（使用相同的过期时间）
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存更新失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
	} else {
		log.Printf("[缓存更新成功] user_id=%d", user.ID)
	}

	return nil
}

// DeleteUser 删除用户
func (s *userServiceWithPenetration) DeleteUser(id int64) error {
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

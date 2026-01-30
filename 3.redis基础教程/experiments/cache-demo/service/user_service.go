package service

import (
	"cache-demo/cache"
	"cache-demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

// userService 用户服务实现
type userService struct {
	repo  model.UserRepo
	cache cache.UserCache
}

// NewUserService 创建用户服务实例
func NewUserService(repo model.UserRepo, cache cache.UserCache) UserService {
	return &userService{
		repo:  repo,
		cache: cache,
	}
}

// GetUserByID 根据ID获取用户（Cache-Aside 模式）
// Cache-Aside 模式流程：
// 1. 先查缓存
// 2. 缓存命中 -> 直接返回
// 3. 缓存未命中 -> 查数据库 -> 写入缓存 -> 返回
func (s *userService) GetUserByID(id int64) (*model.User, error) {
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
		// 其他数据库错误（如连接失败等）
		log.Printf("[数据库查询失败] user_id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 如果用户不存在，返回nil（注意：此实现不缓存空值，存在缓存穿透风险）
	// 如需防止缓存穿透，请使用 user_service_penetration.go 中的实现
	if user == nil {
		log.Printf("[用户不存在] user_id=%d", id)
		return nil, fmt.Errorf("用户不存在: user_id=%d", id)
	}

	// 3. 写入缓存（设置过期时间 5 分钟）
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
		// 缓存写入失败不影响业务逻辑，只记录日志
	} else {
		log.Printf("[缓存写入成功] user_id=%d, username=%s, expire=%d秒", id, user.Username, cache.DefaultExpireSeconds)
	}

	return user, nil
}

// CreateUser 创建用户
// 创建用户时不需要更新缓存（新用户，缓存中不存在）
func (s *userService) CreateUser(user *model.User) error {
	log.Printf("[创建用户] username=%s", user.Username)

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	log.Printf("[创建用户成功] user_id=%d, username=%s", user.ID, user.Username)
	return nil
}

// UpdateUser 更新用户
// 更新用户时，需要同时更新缓存
func (s *userService) UpdateUser(user *model.User) error {
	log.Printf("[更新用户] user_id=%d", user.ID)

	// 1. 更新数据库
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	// 2. 更新缓存（使用相同的过期时间）
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存更新失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
		// 缓存更新失败不影响业务逻辑
	} else {
		log.Printf("[缓存更新成功] user_id=%d", user.ID)
	}

	return nil
}

// DeleteUser 删除用户
// 删除用户时，需要同时删除缓存
func (s *userService) DeleteUser(id int64) error {
	log.Printf("[删除用户] user_id=%d", id)

	// 1. 删除数据库记录
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	// 2. 删除缓存
	if err := s.cache.DeleteUser(id); err != nil {
		log.Printf("[缓存删除失败] user_id=%d, error=%v (不影响业务逻辑)", id, err)
		// 缓存删除失败不影响业务逻辑
	} else {
		log.Printf("[缓存删除成功] user_id=%d", id)
	}

	return nil
}

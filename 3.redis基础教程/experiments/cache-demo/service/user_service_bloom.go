package service

import (
	"cache-demo/cache"
	"cache-demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// userServiceWithBloom 支持布隆过滤器的用户服务实现
type userServiceWithBloom struct {
	repo  model.UserRepo
	cache cache.UserCacheWithBloom
}

// NewUserServiceWithBloom 创建支持布隆过滤器的用户服务实例
func NewUserServiceWithBloom(repo model.UserRepo, cache cache.UserCacheWithBloom) UserService {
	return &userServiceWithBloom{
		repo:  repo,
		cache: cache,
	}
}

// GetUserByID 根据ID获取用户（使用布隆过滤器防止缓存穿透）
func (s *userServiceWithBloom) GetUserByID(id int64) (*model.User, error) {
	// 1. 先检查布隆过滤器（第一道防线）
	exists, err := s.cache.ExistsInBloomFilter(id)
	if err != nil {
		log.Printf("[布隆过滤器查询失败] user_id=%d, error=%v (继续查询)", id, err)
		// 布隆过滤器查询失败，继续查询（容错处理）
	} else if !exists {
		// 布隆过滤器判断不存在，一定不存在，直接返回
		log.Printf("[布隆过滤器拦截] user_id=%d 不存在 (防止缓存穿透)", id)
		return nil, fmt.Errorf("用户不存在: user_id=%d", id)
	}

	// 2. 布隆过滤器判断可能存在，继续查询缓存
	user, err := s.cache.GetUser(id)
	if err == nil && user != nil {
		log.Printf("[缓存命中] user_id=%d, username=%s", id, user.Username)
		return user, nil
	}

	// 3. 缓存未命中，查数据库
	log.Printf("[缓存未命中] user_id=%d, 查询数据库", id)
	user, err = s.repo.FindByID(id)
	if err != nil {
		// 检查是否是记录不存在的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[用户不存在] user_id=%d (布隆过滤器误判)", id)
			// 注意：这里不添加到布隆过滤器，因为数据不存在
			// 如果添加会导致误判率增加
			return nil, fmt.Errorf("用户不存在: user_id=%d", id)
		}
		// 其他数据库错误
		log.Printf("[数据库查询失败] user_id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 4. 用户存在，写入缓存和布隆过滤器
	log.Printf("[用户存在] user_id=%d, username=%s, 写入缓存和布隆过滤器", id, user.Username)
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存写入失败] user_id=%d, error=%v (不影响返回结果)", id, err)
	} else {
		log.Printf("[缓存写入成功] user_id=%d, username=%s, expire=%d秒", id, user.Username, cache.DefaultExpireSeconds)
	}

	// 添加到布隆过滤器（只添加存在的用户）
	if err := s.cache.AddToBloomFilter(id); err != nil {
		log.Printf("[布隆过滤器添加失败] user_id=%d, error=%v (不影响返回结果)", id, err)
	} else {
		log.Printf("[布隆过滤器添加成功] user_id=%d", id)
	}

	return user, nil
}

// CreateUser 创建用户
func (s *userServiceWithBloom) CreateUser(user *model.User) error {
	log.Printf("[创建用户] username=%s", user.Username)

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	// 添加到布隆过滤器
	if err := s.cache.AddToBloomFilter(user.ID); err != nil {
		log.Printf("[布隆过滤器添加失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
	} else {
		log.Printf("[布隆过滤器添加成功] user_id=%d", user.ID)
	}

	log.Printf("[创建用户成功] user_id=%d, username=%s", user.ID, user.Username)
	return nil
}

// UpdateUser 更新用户
func (s *userServiceWithBloom) UpdateUser(user *model.User) error {
	log.Printf("[更新用户] user_id=%d", user.ID)

	// 1. 更新数据库
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	// 2. 更新缓存
	if err := s.cache.SetUser(user, cache.DefaultExpireSeconds); err != nil {
		log.Printf("[缓存更新失败] user_id=%d, error=%v (不影响业务逻辑)", user.ID, err)
	} else {
		log.Printf("[缓存更新成功] user_id=%d", user.ID)
	}

	// 注意：布隆过滤器不支持删除，所以不需要更新
	// 如果用户ID不变，布隆过滤器中已经存在

	return nil
}

// DeleteUser 删除用户
func (s *userServiceWithBloom) DeleteUser(id int64) error {
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

	// 注意：布隆过滤器不支持删除，所以无法从布隆过滤器中删除
	// 这会导致误判率略微增加，但影响不大

	return nil
}

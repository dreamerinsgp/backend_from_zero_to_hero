package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(50);uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"column:email;type:varchar(100);index;not null" json:"email"`
	Age       int       `gorm:"column:age;type:int;default:0" json:"age"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// UserRepo 用户仓储接口
type UserRepo interface {
	FindByID(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

// userRepo 用户仓储实现
type userRepo struct {
	db *gorm.DB
}

// NewUserRepo 创建用户仓储实例
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

// FindByID 根据ID查询用户
func (r *userRepo) FindByID(id int64) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername 根据用户名查询用户
func (r *userRepo) FindByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建用户
func (r *userRepo) Create(user *User) error {
	return r.db.Create(user).Error
}

// Update 更新用户
func (r *userRepo) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *userRepo) Delete(id int64) error {
	return r.db.Delete(&User{}, id).Error
}

package echoapp

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type RegisterParam struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type LoginParam struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"mobile"`
	Pwd    string `json:"password"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func (r *RegisterUser) TableName() string {
	return "users"
}

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Phone  string `json:"mobile"`
	Score  int    `json:"score"`
	Role   string `json:"role"`
}

type UserScoreParam struct {
	UserId int `json:"user_id"`
	Score  int `json:"score"`
}

type Role struct {
	gorm.Model
	ID        uint   `gorm:"AUTO_INCREMENT"`
	Key       string `gorm:"size:255"`
	Name      string `gorm:"size:255；unique"`
	GuardName string `json:"guard_name;not null"`
}

type Permission struct {
	gorm.Model
	ID        uint
	Name      string
	GuardName string `json:"guard_name"`
}
type RoleandPermissionParam struct {
	Role       string `json:"role"`
	Permission string `json:"permission"`
}
type Role_Has_Permission struct {
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}

type UserService interface {
	Save(ctx echo.Context, user *User) error
	AddScore(ctx echo.Context, user *User, amount int) error
	SubScore(ctx echo.Context, user *User, amount int) error
	Login(ctx echo.Context, param *LoginParam) (*User, error)
	GetUserById(ctx echo.Context, userId int) (*User, error)
	Create(c echo.Context, user *RegisterUser) error
	Addroles(c echo.Context, param *Role) error
	AddPermission(c echo.Context, param *Permission) error
	RegisterUser(c echo.Context, param *RegisterUser) error
	RoleHasPermission(c echo.Context, param *RoleandPermissionParam) (*Role_Has_Permission, error)
}
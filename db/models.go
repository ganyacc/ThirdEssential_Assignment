package db

import "time"

type Users struct {
	Id       int `gorm:"primary_key"`
	Name     string
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"not_null" json:"-"`
	Role     string
	Phone    string
	Address  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Product struct {
	Id       int `gorm:"priamary_key"`
	Name     string
	ImageUrl *string `gorm:"type:text"`
	Price    float32
	UserId   int `gorm:"not_null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserProdActivity struct {
	Id        int `gorm:"primary_key"`
	ProductID int
	UserID    int
	Type      string
	Image     string
	ProdName  string
	Price     float32

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLoginActivity struct {
	Id         int `gorm:"primary_key"`
	UserID     int
	LogintTime time.Time
	LogOutTime time.Time

	CreatedAt time.Time
	UpdateAt  time.Time
}

type Token struct {
	Id     int `gorm:"primary_key"`
	Token  string
	Expiry time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID         uint      `json:"id" gorm:"PrimaryKey"`
	UserName   string    `json:"username"`
	Email      string    `json:"email" gorm:"index;unique; not null;"`
	Password   string    `json:"password"`
	FirtsName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Expiry     time.Time `json:"expiry"`
	Verified   bool      `json:"verified" gorm:"default:false"`
	UserType   string    `json:"user_type" gorm:"default:buyer"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdateddAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}

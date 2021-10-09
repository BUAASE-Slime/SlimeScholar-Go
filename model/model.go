package model

// 用户
type User struct {
	UserID       uint64 `gorm:"primary_key; not null;" json:"user_id"`
	Username     string `gorm:"size:25; not null; unique" json:"username"`
	Password     string `gorm:"size:25; not null" json:"password"`
	UserInfo     string `gorm:"size:255;" json:"user_info"`
	UserType     uint64 `gorm:"default:0" json:"user_type"` // 0: 普通用户，1: 认证机构用户
	Affiliation  string `gorm:"size:25;" json:"affiliation"`
	Email        string `gorm:"size:50;" json:"email"`
	HasComfirmed bool   `gorm:"default:false" json:"has_comfirmed"`
	ConfirmNumber int `gorm:"default:0" json:"confirm_number"`
}

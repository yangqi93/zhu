package models

type Topic struct {
	ID          int64  `gorm:"primary_key;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Image       string `gorm:"column:image"`
	Status      int32  `gorm:"column:status"` // 0：禁用 1：启用
	CreatedAt   int64  `gorm:"column:created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at"`
}

func (Topic) TableName() string {
	return "topic"
}

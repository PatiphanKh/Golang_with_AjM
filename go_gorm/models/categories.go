package models

type Category struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"type:longtext"`
	Description string         `gorm:"type:longtext"`
}

func (Category) TableName() string {
	return "categories"
}

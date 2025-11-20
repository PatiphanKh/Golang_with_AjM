package models

type Product struct {
	ID          uint64  `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int64   `gorm:"not null;default:0"`
	CategoryID  *uint64 `gorm:"index"`

	Category Category `gorm:"foreignKey:Category;references:ID"`
}

func (Product) TableName() string {
	return "products"
}

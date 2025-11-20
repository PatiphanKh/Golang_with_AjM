package models

type OrderItem struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	OrderID   *uint64 `gorm:"index"`
	ProductID *uint64 `gorm:"index"`

	Quantity *int64
	Price    *float64

	Product Product `gorm:"foreignKey:ProductID"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

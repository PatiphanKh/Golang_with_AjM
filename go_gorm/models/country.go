package models

type Country struct {
    Idx  int    `gorm:"primaryKey;autoIncrement"`
    Name string `gorm:"type:longtext;not null"`
}

func (Country) TableName() string {
    return "country"
}

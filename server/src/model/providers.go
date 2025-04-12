package model

type Provider struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:text;not null" json:"name"`
}

package model

import (
	"time"
)

type History struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProviderID  uint      `gorm:"not null" json:"provider_id"`
	TokenID     uint      `gorm:"not null" json:"token_id"`
	OperationID uint      `gorm:"not null" json:"operation_id"`
	APR         float32   `gorm:"type:real;not null" json:"apr"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`

	Provider  Provider  `gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"provider"`
	Token     Token     `gorm:"foreignKey:TokenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"token"`
	Operation Operation `gorm:"foreignKey:OperationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"operation"`
}

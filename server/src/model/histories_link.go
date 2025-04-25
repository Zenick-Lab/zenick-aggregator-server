package model

type HistoryLink struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProviderID  uint   `gorm:"not null" json:"provider_id"`
	TokenID     uint   `gorm:"not null" json:"token_id"`
	OperationID uint   `gorm:"not null" json:"operation_id"`
	Link        string `gorm:"type:text;not null" json:"link"`

	Provider  Provider  `gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"provider"`
	Token     Token     `gorm:"foreignKey:TokenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"token"`
	Operation Operation `gorm:"foreignKey:OperationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"operation"`
}

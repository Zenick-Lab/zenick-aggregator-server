package model

type LiquidityPoolHistoryLink struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProviderID uint   `gorm:"not null" json:"provider_id"`
	TokenAID   uint   `gorm:"not null" json:"token_a_id"`
	TokenBID   uint   `gorm:"not null" json:"token_b_id"`
	Link       string `gorm:"type:text;not null" json:"link"`

	Provider Provider `gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"provider"`
	TokenA   Token    `gorm:"foreignKey:TokenAID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"token_a"`
	TokenB   Token    `gorm:"foreignKey:TokenBID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"token_b"`
}

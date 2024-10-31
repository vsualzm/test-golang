package model

type Product struct {
	ID           uint    `gorm:"primaryKey"`
	SKU          string  `gorm:"unique;not null"`
	Name         string  `gorm:"not null"`
	Price        float64 `gorm:"not null"`
	InventoryQty int     `gorm:"not null"`
}

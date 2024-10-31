package model

type Cart struct {
	ID    int `gorm:"primaryKey"`
	Items []CartItem
}

type CartItem struct {
	ID        int `gorm:"primaryKey"`
	CartID    int
	ProductID int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}

package service

import (
	"errors"
	"test-golang/model"

	"gorm.io/gorm"
)

type CartService struct {
	DB *gorm.DB
}

func (s *CartService) AddToCart(cartID int, productID int, quantity int) error {
	var product model.Product
	if err := s.DB.First(&product, productID).Error; err != nil {
		return errors.New("product not found")
	}
	if product.InventoryQty < quantity {
		return errors.New("insufficient product inventory")
	}

	var cart model.Cart
	if err := s.DB.FirstOrCreate(&cart, model.Cart{ID: cartID}).Error; err != nil {
		return err
	}

	cartItem := model.CartItem{
		CartID:    cart.ID,
		ProductID: int(product.ID),
		Quantity:  quantity,
	}
	return s.DB.Create(&cartItem).Error
}

func (s *CartService) ViewCart(cartID int) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	if err := s.DB.Preload("Product").Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (s *CartService) Checkout(cartID int) (float64, error) {
	var cartItems []model.CartItem
	if err := s.DB.Preload("Product").Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return 0, err
	}

	var total float64
	var macbookProCount, googleHomeCount, alexaSpeakerCount int

	// Hitung jumlah item untuk promosi
	for _, item := range cartItems {
		switch item.Product.SKU {
		case "43N23P": // MacBook Pro
			macbookProCount += item.Quantity
		case "120P90": // Google Home
			googleHomeCount += item.Quantity
		case "A304SD": // Alexa Speaker
			alexaSpeakerCount += item.Quantity
		}
	}

	for _, item := range cartItems {
		productPrice := item.Product.Price
		quantity := item.Quantity

		switch item.Product.SKU {
		case "43N23P": // MacBook Pro
			total += float64(quantity) * productPrice
			if macbookProCount > 0 {
				// Berikan Raspberry Pi B gratis per MacBook Pro
				var raspberryPi model.Product
				if err := s.DB.Where("sku = ?", "234234").First(&raspberryPi).Error; err == nil && raspberryPi.InventoryQty > 0 {
					s.DB.Model(&raspberryPi).Update("inventory_qty", raspberryPi.InventoryQty-1)
				}
			}

		case "120P90": // Google Home
			if googleHomeCount >= 3 {
				// Buy 3 for the price of 2
				total += float64((quantity/3*2)+(quantity%3)) * productPrice
			} else {
				total += float64(quantity) * productPrice
			}

		case "A304SD": // Alexa Speaker
			if alexaSpeakerCount >= 4 {
				// Diskon 10% jika membeli 4 atau lebih
				total += float64(quantity) * productPrice * 0.9
			} else {
				total += float64(quantity) * productPrice
			}

		default:
			// Produk tanpa promosi
			total += float64(quantity) * productPrice
		}

		// Kurangi stok produk dari database
		if item.Product.InventoryQty < quantity {
			return 0, errors.New("insufficient inventory for product " + item.Product.Name)
		}
		s.DB.Model(&item.Product).Update("inventory_qty", item.Product.InventoryQty-item.Quantity)
	}

	// Hapus item dari cart setelah checkout
	if err := s.DB.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error; err != nil {
		return 0, err
	}

	return total, nil
}

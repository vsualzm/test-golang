package handler

// type CartHandler struct {
// 	CartService *service.CartService
// }

// // AddToCart handler untuk menambahkan item ke keranjang
// func (h *CartHandler) AddToCart(c *gin.Context) {
// 	cartID, _ := strconv.Atoi(c.Param("cart_id"))
// 	var input struct {
// 		ProductID int `json:"product_id"`
// 		Quantity  int `json:"quantity"`
// 	}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	if err := h.CartService.AddToCart(cartID, input.ProductID, input.Quantity); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
// }

// // ViewCart handler untuk melihat isi keranjang
// func (h *CartHandler) ViewCart(c *gin.Context) {
// 	cartID, _ := strconv.Atoi(c.Param("cart_id"))

// 	items, err := h.CartService.ViewCart(cartID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"cart_items": items})
// }

// // Checkout handler untuk menyelesaikan pembelian dan mengurangi stok produk
// func (h *CartHandler) Checkout(c *gin.Context) {
// 	cartID, _ := strconv.Atoi(c.Param("cart_id"))

// 	total, err := h.CartService.Checkout(cartID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"total": total, "message": "Checkout successful"})
// }

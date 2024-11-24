package modal

// CartItem represents an item in the cart.
type CartItem struct {
	CartID          int     `json:"cart_id"`
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float64 `json:"product_price"`
	Quantity        int     `json:"quantity"`
	TotalPrice      float64 `json:"total_price"`
	DateCreated     string  `json:"date_created"`
	DateUpdated     string  `json:"date_updated"`
	Store           string  `json:"store"`
	Image           string  `json:"image"`
	Brand           string  `json:"brand"`
	PromotionName   string  `json:"promotion_name"`
	DiscountMoney   int     `json:"discount_money"`
	DiscountPercent int     `json:"discount_percent"`
}

// Cart represents a cart model.
type Cart struct {
	ProductID   int     `json:"product_id"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	Store       string  `json:"store"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
}

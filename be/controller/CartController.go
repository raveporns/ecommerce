package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raveporns/ecommerce/modals" 
)

// CartController struct holds the DB connection.
type CartController struct {
	DB *sql.DB
}

// getCartByQuery is a helper function to query cart items based on a SQL query and parameters.
func (cc *CartController) getCartByQuery(query string, args ...interface{}) ([]modal.CartItem, error) {
	rows, err := cc.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []modal.CartItem

	// Loop through the rows and scan each record
	for rows.Next() {
		var cartItem modal.CartItem
		if err := rows.Scan(&cartItem.CartID, &cartItem.ProductID, &cartItem.ProductName, &cartItem.ProductPrice,
			&cartItem.Quantity, &cartItem.TotalPrice, &cartItem.DateCreated, &cartItem.DateUpdated,
			&cartItem.Store, &cartItem.Image, &cartItem.Brand , &cartItem.PromotionName,
			&cartItem.DiscountMoney,
			&cartItem.DiscountPercent,); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return cartItems, nil
}

// GetCartAll returns all cart items.
func (cc *CartController) GetCartAll(c *gin.Context) {

	query := `SELECT c.cart_id, c.product_id, p.product_name, p.product_price, c.quantity, c.total_price, 
		c.date_created, c.date_updated, c.store , p.image , p.brand ,p.promotion_name, p.discount_money, p.discount_percent
		FROM cart c JOIN product p on c.product_id = p.product_id
		`

	cartItems, err := cc.getCartByQuery(query)
	if err != nil {
		log.Printf("Error fetching cart for user : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cart"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

// AddProductToCart adds a product to the user's cart.
func (cc *CartController) AddProductToCart(c *gin.Context) {
	var cart modal.Cart

	// Bind JSON request to cart object
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate required fields
	if cart.ProductID == 0 || cart.Quantity == 0 || cart.Store == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID, Quantity, and Store are required fields"})
		return
	}

	// Calculate the total price
	totalPrice := float64(cart.Quantity) * cart.TotalPrice

	// Prepare the SQL query to insert the cart item
	query := `INSERT INTO cart (product_id, quantity, total_price, store, date_created, date_updated)
        VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := cc.DB.Exec(query, cart.ProductID, cart.Quantity, totalPrice, cart.Store, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error adding product to cart: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding product to cart"})
		return
	}

	// Return success response with added product details
	c.JSON(http.StatusOK, gin.H{
		"message":     "Product added to cart",
		"product_id":  cart.ProductID,
		"quantity":    cart.Quantity,
		"total_price": totalPrice,
		"store":       cart.Store,
	})
}

// UpdateProductQuantity updates the quantity of a product in the cart.
func (cc *CartController) UpdateProductQuantity(c *gin.Context) {
	cartID := c.Param("cart_id")
	var cart modal.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Update the quantity and total price of the product in the cart
	query := `UPDATE cart SET quantity = $1, total_price = $2, date_updated = $3 WHERE cart_id = $4`
	_, err := cc.DB.Exec(query, cart.Quantity, cart.TotalPrice, time.Now(), cartID)
	if err != nil {
		log.Printf("Error updating cart item: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product quantity updated"})
}

// RemoveProductFromCart removes a product from the user's cart.
func (cc *CartController) RemoveProductFromCart(c *gin.Context) {
	cartID := c.Param("cart_id")

	// Delete the product from the cart
	query := `DELETE FROM cart WHERE cart_id = $1`
	_, err := cc.DB.Exec(query, cartID)
	if err != nil {
		log.Printf("Error removing product from cart: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing product from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}

package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	modal "github.com/raveporns/ecommerce/modals" // ใช้ models แทน modals
)

type ProductController struct {
	DB *sql.DB
}

// GetProductAll ดึงข้อมูลสินค้าทั้งหมดจากฐานข้อมูล
func (pc *ProductController) GetProductAll(c *gin.Context) {
	query := `SELECT 
    p.product_id, 
    p.product_name, 
    p.product_price, 
    p.product_quality, 
    p.date_create, 
    p.date_update, 
    p.product_type, 
    p.brand, 
    p.store, 
    p.status, 
    p.product_description, 
    p.image, 
    p.promotion_name,
	p.discount_money,
	p.discount_percent
	FROM product p
	`
	rows, err := pc.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID ดึงข้อมูลสินค้าจากฐานข้อมูลตาม ID
func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
        WHERE product_id = $1
    `
	row := pc.DB.QueryRow(query, id)

	product, err := modal.ScanProductRow(row) // ใช้ ScanProductRow จาก models
	if err != nil {
		log.Printf("Error fetching product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductBrand ดึงข้อมูลสินค้าจากฐานข้อมูลตามแบรนด์
func (pc *ProductController) GetProductBrand(c *gin.Context) {
	brand := c.Param("brand")
	if brand == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand parameter is required"})
		return
	}

	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
        WHERE brand = $1
    `
	rows, err := pc.DB.Query(query, brand)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductPrice ดึงข้อมูลสินค้าจัดเรียงตามราคาสินค้า
func (pc *ProductController) GetProductPrice(c *gin.Context) {
	order := c.DefaultQuery("order", "asc")
	var query string
	if order == "desc" {
		query = `
            SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
            ORDER BY product_price DESC
        `
	} else {
		query = `
            SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
            ORDER BY product_price ASC
        `
	}

	rows, err := pc.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductStore(c *gin.Context) {
	store := c.Param("store")
	if store == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Store parameter is required"})
		return
	}

	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
        WHERE store = $1
    `
	rows, err := pc.DB.Query(query, store)
	if err != nil {
		log.Printf("Error fetching products from store: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products from store"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product) // เพิ่ม product ลงใน products
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products) // ส่งคืนข้อมูลสินค้าทั้งหมดที่เกี่ยวข้องกับ store
}

func (pc *ProductController) GetProductStoreFromEachBrand(c *gin.Context) {
	store := c.Param("store")
	brand := c.Param("brand")

	// Check if store and brand parameters are present
	if store == "" || brand == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Store and brand parameters are required"})
		return
	}

	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money, discount_percent
        FROM product
        WHERE store = $1 AND brand = $2
    `
	// Use both store and brand as query parameters
	rows, err := pc.DB.Query(query, store, brand)
	if err != nil {
		log.Printf("Error fetching products from store: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products from store"})
		return
	}
	defer rows.Close()

	var products []modal.Product // Use the correct model (modal.Product)

	// Iterate over the rows and scan into products
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // Ensure ScanProduct function works properly
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product) // Append the scanned product to the products slice
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	// Return the fetched products as JSON
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductStatus(c *gin.Context) {
	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status parameter is required"})
		return
	}

	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
        WHERE status = $1
		LIMIT 3
    `
	rows, err := pc.DB.Query(query, status)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product) // เพิ่ม product ลงใน products
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products) // ส่งข้อมูลสินค้าทั้งหมดที่มีสถานะตามที่ระบุ
}

func (pc *ProductController) GetNewProductsFromEachStore(c *gin.Context) {
	query := `
        SELECT DISTINCT ON (store) 
            product_id, 
            product_name, 
            product_price, 
            product_quality, 
            date_create, 
            date_update, 
            product_type, 
            brand, 
            store, 
            status, 
            product_description, 
            image, promotion_name, 
			discount_money,
			discount_percent
        FROM product
        ORDER BY store, date_create DESC
    `

	rows, err := pc.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductDiscountMoney(c *gin.Context) {
	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name, discount_money,discount_percent
        FROM product
		where discount_money > 0
        order by discount_money desc
        limit 3
    `
	rows, err := pc.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product) // เพิ่ม product ลงใน products
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products) // ส่งข้อมูลสินค้าทั้งหมดที่มีสถานะตามที่ระบุ
}

func (pc *ProductController) GetProductDiscountPercent(c *gin.Context) {

	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, image, promotion_name,discount_money, discount_percent
        FROM product
		where discount_percent > 0
        order by discount_percent desc
        limit 3
    `
	rows, err := pc.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product) // เพิ่ม product ลงใน products
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, products) // ส่งข้อมูลสินค้าทั้งหมดที่มีสถานะตามที่ระบุ
}

func (pc *ProductController) GetProductBySearch(c *gin.Context) {
	search := c.DefaultQuery("search", "") // รับค่าคำค้นหาจาก query string
	if search == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search parameter is required"})
		return
	}

	// แก้ไข query ให้ค้นหาจาก product_name และ brand เท่านั้น
	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, 
               product_type, brand, store, status, product_description, image, 
               promotion_name, discount_money, discount_percent
        FROM product
        WHERE product_name ILIKE $1 OR brand ILIKE $1
    `

	// ใช้ query สำหรับการค้นหาผ่านคำค้นหาจาก query string
	rows, err := pc.DB.Query(query, "%_"+search+"_%") // เพิ่ม '%' เพื่อให้ค้นหาจากคำค้นหาผสมใน product_name หรือ brand
	if err != nil {
		log.Printf("Error fetching products: %v", err) // เพิ่ม log เพื่อให้เห็นข้อผิดพลาด
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct สำหรับแปลงข้อมูล
		if err != nil {
			log.Printf("Error scanning row: %v", err) // เพิ่ม log เพื่อให้เห็นข้อผิดพลาด
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product) // เพิ่มสินค้าที่ถูกแปลงไปใน slice
	}

	// ตรวจสอบ error หลังจากวน loop
	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err) // เพิ่ม log เพื่อตรวจสอบข้อผิดพลาด
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	// ตรวจสอบว่ามีผลลัพธ์หรือไม่
	if len(products) == 0 {
		log.Printf("No products found for search term: %v", search) // เพิ่ม log ถ้าไม่พบผลลัพธ์
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}

	// ส่งข้อมูลที่ได้ไปยัง client
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductSearch(c *gin.Context) {
	// รับค่าคำค้นหาจาก query parameter
	search := c.DefaultQuery("query", "") // เปลี่ยนเป็นใช้ query string แทน path parameter
	if search == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกคำค้นหา"})
		return
	}

	// การค้นหาด้วย ILIKE สำหรับคำค้นหาที่ไม่ตรงเป๊ะ
	query := `
        SELECT product_id, product_name, product_price, product_quality, date_create, date_update, 
               product_type, brand, store, status, product_description, image, 
               promotion_name, discount_money, discount_percent
        FROM product
        WHERE product_name ILIKE '%' || $1 || '%' OR brand ILIKE '%' || $1 || '%'
        LIMIT 100
    `

	// ค้นหาผลลัพธ์จากฐานข้อมูล
	rows, err := pc.DB.Query(query, search)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []modal.Product // ใช้ models แทน modal
	for rows.Next() {
		product, err := modal.ScanProduct(rows) // ใช้ ScanProduct จาก models
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	// ส่งผลลัพธ์การค้นหากลับไปยังผู้ใช้
	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "ไม่พบผลิตภัณฑ์ที่ตรงกับคำค้นหา"})
	} else {
		c.JSON(http.StatusOK, products)
	}
}

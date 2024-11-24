package modal

import "database/sql"

// Product คือตัวแทนของสินค้าหนึ่งตัวในระบบ
type Product struct {
	ProductID          int    `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductPrice       int    `json:"product_price"`
	ProductQuality     int    `json:"product_quality"`
	DateCreate         string `json:"date_create"`
	DateUpdate         string `json:"date_update"`
	ProductType        string `json:"product_type"`
	Brand              string `json:"brand"`
	Store              string `json:"store"`
	Status             string `json:"status"`
	ProductDescription string `json:"product_description"`
	Image              string `json:"image"`
	PromotionName      string `json:"promotion_name"`
	DiscountMoney      int    `json:"discount_money"`
	DiscountPercent    int    `json:"discount_percent"`
}

// ScanProduct ใช้สำหรับแปลงข้อมูลจากฐานข้อมูลเป็น Product struct
func ScanProduct(rows *sql.Rows) (*Product, error) {
	var product Product
	var image sql.NullString

	err := rows.Scan(
		&product.ProductID,
		&product.ProductName,
		&product.ProductPrice,
		&product.ProductQuality,
		&product.DateCreate,
		&product.DateUpdate,
		&product.ProductType,
		&product.Brand,
		&product.Store,
		&product.Status,
		&product.ProductDescription,
		&image,
		&product.PromotionName,
		&product.DiscountMoney,
		&product.DiscountPercent,
	)
	if err != nil {
		return nil, err
	}

	// แปลง image ที่อาจจะเป็น NULL
	if image.Valid {
		product.Image = image.String
	} else {
		product.Image = "" // กำหนดค่าเป็นค่าว่างถ้าไม่มีภาพ
	}

	return &product, nil
}

func ScanProductRow(row *sql.Row) (*Product, error) {
	var product Product
	var image sql.NullString

	err := row.Scan(
		&product.ProductID,
		&product.ProductName,
		&product.ProductPrice,
		&product.ProductQuality,
		&product.DateCreate,
		&product.DateUpdate,
		&product.ProductType,
		&product.Brand,
		&product.Store,
		&product.Status,
		&product.ProductDescription,
		&image,
		&product.PromotionName,
		&product.DiscountMoney,
		&product.DiscountPercent,
	)
	if err != nil {
		return nil, err
	}

	// แปลง image ที่อาจจะเป็น NULL
	if image.Valid {
		product.Image = image.String
	} else {
		product.Image = "" // กำหนดค่าเป็นค่าว่างถ้าไม่มีภาพ
	}

	return &product, nil
}

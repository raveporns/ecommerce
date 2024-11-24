package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/raveporns/ecommerce/controller"
)

func main() {
	// โหลด environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึงค่าจาก environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	// สร้าง DSN สำหรับเชื่อมต่อกับฐานข้อมูล PostgreSQL
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	// เชื่อมต่อกับฐานข้อมูล
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อฐานข้อมูล
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	ProductController := &controller.ProductController{DB: db}
	CartController := &controller.CartController{DB: db}
	ReviewController := &controller.ReviewController{DB: db}

	// สร้าง Gin router
	router := gin.Default()

	// Middleware สำหรับจัดการ CORS
	router.Use(cors.Default())

	router.GET("/products", ProductController.GetProductAll)
	// router.GET("/products/brand/:brand", ProductController.GetProductBrand)
	router.GET("/products/status/:status", ProductController.GetProductStatus)
	router.GET("/products/:id", ProductController.GetProductByID)
	router.GET("/brand/:brand", ProductController.GetProductBrand)
	router.GET("/store/:store", ProductController.GetProductStore)
	router.GET("/store/:store/:brand", ProductController.GetProductStoreFromEachBrand)

	router.GET("/products/price/:price", ProductController.GetProductPrice)
	router.GET("/search", ProductController.GetProductSearch)
	router.GET("/products/new", ProductController.GetNewProductsFromEachStore)

	router.GET("/discount/money", ProductController.GetProductDiscountMoney)
	router.GET("/discount/percent", ProductController.GetProductDiscountPercent)

	commentRoutes := router.Group("/review")
	{
		commentRoutes.GET("", ReviewController.GetCommentAll)
		commentRoutes.POST("/add", ReviewController.AddComment)
		commentRoutes.PUT("/update", ReviewController.UpdateComment)
		commentRoutes.DELETE("/remove", ReviewController.RemoveComment)

	}

	cartRoutes := router.Group("/cart")
	{

		// cartRoutes.GET("/user/:user_id", CartController.GetCartByUserID) // ดึงข้อมูลตะกร้าของผู้ใช้
		cartRoutes.GET("", CartController.GetCartAll)                               // ดึงข้อมูลตะกร้าของผู้ใช้
		cartRoutes.POST("/add", CartController.AddProductToCart)                    // เพิ่มสินค้าลงในตะกร้า
		cartRoutes.PUT("/update/:cart_id", CartController.UpdateProductQuantity)    // อัปเดตจำนวนสินค้าในตะกร้า
		cartRoutes.DELETE("/remove/:cart_id", CartController.RemoveProductFromCart) // ลบสินค้าออกจากตะกร้า
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	log.Println("Server starting on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

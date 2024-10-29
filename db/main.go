package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/joho/godotenv"
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

	// สร้าง Gin router
	router := gin.Default()

	// Middleware สำหรับจัดการ CORS
	router.Use(cors.Default())

	// ตัวอย่างการเพิ่ม route
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

package tests

import (
	"log"
	"time"

	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=1010970549 dbname=healthmate_auth_test_service port=2025 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Không kết nối được DB test: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Lỗi khi lấy sqlDB: %v", err)
	}
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	if err := db.AutoMigrate(&user.Users{}); err != nil {
		log.Fatalf("AutoMigrate lỗi: %v", err)
	}

	log.Println("✅ Kết nối DB test thành công và đã migrate bảng users")
	return db
}
package config

import (
	"courier-api/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(
		&models.User{},
		&models.Courier{},
		&models.Tariff{},
		&models.Shipment{},
		&models.ShipmentItem{},
		&models.TrackingHistory{},
	)

	DB = database
	fmt.Println("Database connected successfully")
}

// CREATE TABLE `shipment_items` (
//   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   `shipment_id` bigint unsigned DEFAULT NULL,
//   `item_name` longtext,
//   `quantity` int DEFAULT NULL,
//   `weight` double DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `idx_shipment_items_deleted_at` (`deleted_at`),
//   KEY `fk_shipment_items_shipment` (`shipment_id`),
//   CONSTRAINT `fk_shipment_items_shipment` FOREIGN KEY (`shipment_id`) REFERENCES `shipments` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

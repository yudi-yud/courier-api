package main

import (
	"courier-api/config"
	"courier-api/controllers"
	"courier-api/middleware"
	"courier-api/services"
	"log"
	"os"

	_ "courier-api/docs"

	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// @title Courier / Tracking API
// @version 1.0
// @description Ini adalah API untuk sistem manajemen pengiriman ekspedisi.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	services.InitAdmin()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/swagger/*", fiberSwagger.New())
	app.Static("/uploads", "./uploads")

	// Controllers
	authController := controllers.NewAuthController()
	shipmentController := controllers.NewShipmentController()
	courierController := controllers.NewCourierController()
	tariffController := controllers.NewTariffController()
	api := app.Group("/api/v1")

	// Public Routes
	api.Post("/login", authController.Login)
	api.Get("/track/:resi", shipmentController.Track)

	// Protected Routes
	protected := api.Use(middleware.JWTProtected())

	// Auth / User Management
	protected.Post("/register", middleware.Authorize("admin"), authController.Register)

	// Courier Routes (Admin Only)
	protected.Post("/couriers", middleware.Authorize("admin"), courierController.Create)
	protected.Get("/couriers", middleware.Authorize("admin"), courierController.GetAll)

	// TARIF ROUTES
	protected.Post("/tariffs", middleware.Authorize("admin"), tariffController.Create)
	protected.Get("/tariffs", middleware.Authorize("admin"), tariffController.GetAll)

	// Shipment Routes
	protected.Post("/shipments", middleware.Authorize("admin"), shipmentController.Create)
	protected.Get("/shipments", shipmentController.GetAll)
	protected.Get("/shipments/stats", shipmentController.GetStats)
	protected.Post("/shipments/:resi/assign", shipmentController.AssignCourier)
	protected.Patch("/shipments/:resi/status", shipmentController.UpdateStatus)
	protected.Post("/shipments/:resi/pod", shipmentController.UploadPOD)

	protected.Get("/shipments/:resi/pdf", shipmentController.GenerateAirwayBillPDF)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}

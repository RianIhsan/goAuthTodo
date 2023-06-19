package boot

import (
	"goAuthTodo/config"
	"goAuthTodo/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func BootApp() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	config.ConnectDB()
	config.RunMigration()
	port := os.Getenv("PORT")
	app := fiber.New()

	routes.InitRoute(app)
	app.Listen(":" + port)
}

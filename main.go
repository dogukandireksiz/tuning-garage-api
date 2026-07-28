package main

import (
	
	"api/routes"
	"github.com/gofiber/fiber/v3/middleware/logger" // Middleware paketimiz
	
	"log" // hata/bilgi yazdırma kütüphanesidir.

	"github.com/gofiber/fiber/v3" //web sunucu altyapısı.
	"api/database"
)


func main()  {
	
	//Fiber dan önce veritabanına bağlanıyoruz
	database.ConnectDB()
	
	app := fiber.New() //Fiber kütüphanesinden yeni bir sunucu nesnesi yaratılır. app değişkeni artık bizim web sunucumuzdur.

	//Middleware: Gözlemci ekleme
	//Bu satır sayesinde API'ye gelen her istek (hangi IP'den geldi,ne kadar sürdü,hangi statü kodunu döndü) terminale yazdırılacak.
	app.Use(logger.New())

	//Rotaları kurma
	routes.SetupRoutes(app)


	log.Println("Garaj API 3000 portunda çalışıyor...")
	log.Fatal(app.Listen(":3000"))
}


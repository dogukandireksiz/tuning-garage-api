package main

import (
	"api/controllers"
	
	"log" // hata/bilgi yazdırma kütüphanesidir.

	"github.com/gofiber/fiber/v3" //web sunucu altyapısı.
	"api/database"
)


func main()  {
	
	//Fiber dan önce veritabanına bağlanıyoruz
	database.ConnectDB()
	
	app := fiber.New() //Fiber kütüphanesinden yeni bir sunucu nesnesi yaratılır. app değişkeni artık bizim web sunucumuzdur.


	//Araçları listeleme
	app.Get("/api/cars",controllers.GetCars)

	// POST : Garaja yeni araç ekleme
	app.Post("api/cars",controllers.CreateCar)

	// PUT : Garajdaki mevcut bir aracı güncelle
	// (:) ile  başlayan kısımlar dinamik rota (route parameter) olarak adlandırılır 1 de yazılsa 99 da yazılsa çalışır.
	app.Put("api/cars/:id",controllers.UpdateCar)

	// Delete
	app.Delete("api/cars/:id",controllers.DeleteCar)

	log.Println("Garaj API 3000 portunda çalışıyor...")
	log.Fatal(app.Listen(":3000"))
}


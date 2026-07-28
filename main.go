package main

import (
	
	"api/routes"
	"github.com/gofiber/fiber/v3/middleware/logger" // Middleware paketimiz
	
	"log" // hata/bilgi yazdırma kütüphanesidir.
	"os"
	"github.com/joho/godotenv" // Yeni ekledik

	"github.com/gofiber/fiber/v3" //web sunucu altyapısı.
	"api/database"
)


func main()  {
	
	// .env dosyasını sisteme yüklüyoruz
	err := godotenv.Load()
	if err != nil{
		log.Println("Uyarı: .env dosyası bulunamadı, sistem çevre değişkenleri kullanılacak.")
	}

	//Fiber dan önce veritabanına bağlanıyoruz
	database.ConnectDB()
	
	app := fiber.New() //Fiber kütüphanesinden yeni bir sunucu nesnesi yaratılır. app değişkeni artık bizim web sunucumuzdur.

	//Middleware: Gözlemci ekleme
	//Bu satır sayesinde API'ye gelen her istek (hangi IP'den geldi,ne kadar sürdü,hangi statü kodunu döndü) terminale yazdırılacak.
	app.Use(logger.New())

	//Rotaları kurma
	routes.SetupRoutes(app)

	// Port numarasını .env den alıyoruz
	port := os.Getenv("PORT")
	if port == ""{
		port = "3000" // Fallback(B planı)
	} 


	log.Println("Garaj API %s portunda çalışıyor...",port)
	log.Fatal(app.Listen(":" + port))
}


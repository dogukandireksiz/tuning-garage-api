package main

import (
	"fmt"
	"log" // hata/bilgi yazdırma kütüphanesidir.

	"github.com/gofiber/fiber/v3" //web sunucu altyapısı.
)

// Veri modeli
type Car struct{
	ID string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Mods []string `json:"mods"` //["Stage 1 ECU","Custom Jant"]
}

// Geçici veritabanı

var garage = []Car{
	{
	ID : "1",
	Brand : "BMW",
	Model : "M3",
	Mods : []string{"Stage 1 ECU Remap","Spor Süspansiyon"},
	},
}


func main()  {
	
	app := fiber.New() //Fiber kütüphanesinden yeni bir sunucu nesnesi yaratılır. app değişkeni artık bizim web sunucumuzdur.


	//Araçları listeleme
	app.Get("/api/cars",func(c fiber.Ctx) error  {
		//garage listesini JSON formatına çevirip istemciye gönderme
		return c.JSON(garage)
	})

	// POST : Garaja yeni araç ekleme
	app.Post("api/cars",func(c fiber.Ctx) error  {
		// içi boş bir araba nesnesi yaratıyoruz.
		var newCar Car

		// istemciden gelen JSON verisini newCar'ın içine kopyalıyoruz. (Bind)
		// hata çıkarsa (örn: JSON formatı bozuksa) if bloğu çalışır
		// c.Bind().JSON(&newCar): Fiber v3'ün veri bağlama metodudur. İstemciden gelen JSON verisini okur ve bizim Car yapımızla eşleştirir.

		if err := c.Bind().JSON(&newCar); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Gelen veri okunamadı veya fotmat hatali",
			})
		}

		// Henüz gerçek bir veritabanı olmadığı için ID yi manuel veriyoruz
		// Mevcut garaj uzunluğuna 1 ekleyip string e çeviriyoruz
		newCar.ID = fmt.Sprintf("%d",len(garage)+1)

		// Yeni aracı garaj listesine ekliyoruz
		garage = append(garage, newCar)

		// 201 Created statü koduyla birlikte eklenen aracı geri dönüyoruz
		return c.Status(fiber.StatusCreated).JSON(newCar)
		//c.Status(fiber.StatusCreated): Sadece veri dönmek yetmez, HTTP iletişim kurallarına uymalıyız. Başarılı bir ekleme işlemi yapıldığında standart olarak 201 Created kodu dönülür. Hata durumunda ise 400 Bad Request döndük.
		
	})

	log.Println("Garaj API 3000 portunda çalışıyor...")
	log.Fatal(app.Listen(":3000"))
}


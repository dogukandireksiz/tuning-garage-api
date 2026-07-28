package main

import (
	"log"
	

	"github.com/gofiber/fiber/v3"
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
	
	app := fiber.New()

	//Araçları listeleme
	app.Get("/api/cars",func(c fiber.Ctx) error  {
		//garage listesini JSON formatına çevirip istemciye gönderme
		return c.JSON(garage)
	})

	log.Println("Garaj API 3000 portunda çalışıyor...")
	log.Fatal(app.Listen(":3000"))
}


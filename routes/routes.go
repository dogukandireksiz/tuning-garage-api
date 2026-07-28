package routes

import (
	"api/controllers"
	"github.com/gofiber/fiber/v3"
)

// SetupRoutes, uygulamadaki tüm rotaları tanımlayan ana fonksiyondur

func SetupRoutes(app *fiber.App)  {
	
	// Ana API grubunu oluşturuyoruz
	api := app.Group("/api")

	// "/api" altında çalışan "cars" alt grubunu oluituruyoruz
	cars := api.Group("/cars")

	//Artık uzun uzun "/api/cars" yazmaya gerek yok
	cars.Get("/",controllers.GetCars)
	cars.Post("/",controllers.CreateCar)
	cars.Put("/:id",controllers.UpdateCar)
	cars.Delete("/:id",controllers.DeleteCar)
	cars.Post("/:id/mods",controllers.AddModification)
}
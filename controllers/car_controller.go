package controllers

import (
    

    "api/models" // Kendi oluşturduğumuz modeli içeri aktarıyoruz
    "api/database"
    "github.com/gofiber/fiber/v3"
)

func GetCars(c fiber.Ctx) error  {
	var cars []models.Car
	//GORM veritabanındaki tüm kayıtları bulup cars dizisine doldur
	database.DB.Find(&cars)
	
	//cars listesini JSON formatına çevirip istemciye gönderme
	return c.JSON(cars)
}

func CreateCar(c fiber.Ctx) error  {
		// içi boş bir araba nesnesi yaratıyoruz.
		var newCar models.Car

		// istemciden gelen JSON verisini newCar'ın içine kopyalıyoruz. (Bind)
		// hata çıkarsa (örn: JSON formatı bozuksa) if bloğu çalışır
		// c.Bind().JSON(&newCar): Fiber v3'ün veri bağlama metodudur. İstemciden gelen JSON verisini okur ve bizim Car yapımızla eşleştirir.

		if err := c.Bind().JSON(&newCar); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Gelen veri okunamadı veya fotmat hatali",
			})
		}

		database.DB.Create(&newCar)

		// 201 Created statü koduyla birlikte eklenen aracı geri dönüyoruz
		return c.Status(fiber.StatusCreated).JSON(newCar)
		//c.Status(fiber.StatusCreated): Sadece veri dönmek yetmez, HTTP iletişim kurallarına uymalıyız. Başarılı bir ekleme işlemi yapıldığında standart olarak 201 Created kodu dönülür. Hata durumunda ise 400 Bad Request döndük.

}

func UpdateCar(c fiber.Ctx) error  {
		// URL den ":id" kısmını yakalıyoruz
		id := c.Params("id")
		var car models.Car

		//Önce aracı veritabanında arıyoruz
		if result := database.DB.First(&car,id); result.Error != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "Araç bulunamadı.",
			})
		}

	

		// İstemciden gelen güncel veriyi okuyoruz
		var updatedData models.Car
		if err := c.Bind().JSON(&updatedData) ; err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Gelen veri okunamadi",
			})
		}

		//Bulduğumuz aracın özelliklerini güncelliyoruz
		car.Brand = updatedData.Brand
		car.Model = updatedData.Model

		database.DB.Save(&car)

		return c.JSON(&car)
		
}


func DeleteCar(c fiber.Ctx) error  {
		id := c.Params("id")
		var car models.Car

		if result := database.DB.First(&car,id); result.Error != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "Araç bulunamadi.",
			})
		}

		database.DB.Delete(&car)

		return c.JSON(fiber.Map{
			"message":"Araç başariyla silindi.", 
		})

}

//Go'da yazdığın kodları çalıştırıp bir silme (DELETE) işlemi yaptığında çok ilginç bir şey olacak. İstek sana başarıyla silindiğini söyleyecek, GET /api/cars yaptığında araç listede görünmeyecek... Ama aslında o araç veritabanından silinmedi!

//models.Car içine eklediğimiz gorm.Model yapısı sayesinde GORM Soft Delete (Yumuşak Silme) uygular. Yani veritabanında satırı yok etmez, sadece deleted_at (silinme tarihi) sütununu o anki saat ile doldurur. GORM, deleted_at sütunu dolu olan verileri standart aramalarda (Find veya First) otomatik olarak gizler. Bu, yanlışlıkla veri silinmelerine karşı harika bir güvenlik ağıdır.
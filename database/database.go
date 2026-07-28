package database

//Bu dosya, veritabanına bağlanma ve tabloları oluşturma işini üstlenecek:

import (
	"log"
	"api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

// DB, uygulamanın her yerinden veritabanına erişmek için kullanacağımız global değişken.

var DB *gorm.DB

func ConnectDB()  {
	// Çevre değişkeninden veritabanı adını okuyoruz
	dbName := os.Getenv("DB_NAME")

	// Eğer bir şekilde .env okunmazsa ve isim boş gelirse, güvenlik önlemi alıyoruz
	if dbName == ""{
		dbName = "default.db"
	}
	
	var err error
	// "garage.db" adında bir SQLite dostasına bağlanıyoruz
	DB, err = gorm.Open(sqlite.Open(dbName),&gorm.Config{})
	if err != nil{
		log.Fatal("Veritabanına bağlanılamadı! Hata: ",err)
	}

	log.Println("Veritabanı bağlantısı başarılı.")

	// AutoMigrate: Eğer 'cars' tablosu yoksa oluşturur, varsa eksik kolonları ekler.
	err = DB.AutoMigrate(&models.Car{},&models.Modification{})
	
	if err != nil{
		log.Fatal("Tablolar oluşturulamadı! Hata: ",err)
	}
	
	log.Println("Veritabanı tabloları hazır (Auto-Migration tamamlandı.)")
}
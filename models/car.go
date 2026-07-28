package models

import (
	"gorm.io/gorm"
)
// Veri modeli
type Car struct{
	gorm.Model // GORM'un standart alanlarını (ID, CreatedAt, UpdatedAt, DeletedAt) otomatik ekler
	Brand string `json:"brand"`
	Modell string `json:"model"`
	//Mods []string `json:"mods"` //["Stage 1 ECU","Custom Jant"]
	// Not: GORM varsayılan olarak string dizilerini ([]string) basitçe kaydedemez.
    // Bu yüzden şimdilik Mods alanını çıkarıyoruz, ileride ilişkisel veritabanı mantığıyla ekleyeceğiz.
	
	// GORM bu listeyi gördüğünde otomatik olarak bire-çok işilkisi kurar.
	Mods []Modification `json:"mods" gorm:"foreignKey:CarID`
}

//gorm.Model satırı sayesinde artık kendi ID string tanımımıza ihtiyacımız kalmadı. GORM bize otomatik artan bir ID (uint) ve silme/güncelleme tarihleri sağlayacak. Ayrıca Soft Delete (silindi olarak işaretleme) özelliği de otomatik gelecek.

// Yeni tablomuz
type Modification struct{
	gorm.Model
	Name string `json:"name"`
	CarID uint `json:"car_id"` //foreign key
}

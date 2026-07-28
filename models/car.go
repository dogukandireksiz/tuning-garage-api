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
}

//gorm.Model satırı sayesinde artık kendi ID string tanımımıza ihtiyacımız kalmadı. GORM bize otomatik artan bir ID (uint) ve silme/güncelleme tarihleri sağlayacak. Ayrıca Soft Delete (silindi olarak işaretleme) özelliği de otomatik gelecek.


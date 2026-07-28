package dto

// CreateCarRequest: Yeni araç eklenirken dışarıdan beklediğimiz veri şablonu
type CreateCarRequest struct{
	// validate etiketleri ile kurallarımızı koyuyoruz:
	// required: Bu alanın boş geçilmesi yasaktır.
	// min=2: Marka adı en az 2 karakter olmalıdır.
	Brand string `json:"brand" validate:"required,min=2"` 
	Modell string `json:"model" validate:"required" `
}

//Bu yapı sayesinde models.Car içindeki ID veya tarih alanlarını yanlışlıkla dışarıdan müdahaleye açmamış oluyoruz. Sadece almamız gereken Brand ve Model verilerini güvenli bir şekilde alıyoruz.
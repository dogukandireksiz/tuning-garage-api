package models

// Veri modeli
type Car struct{
	ID string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Mods []string `json:"mods"` //["Stage 1 ECU","Custom Jant"]
}
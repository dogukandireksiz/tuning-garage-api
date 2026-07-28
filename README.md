# 🏎️ Tuning Garajı REST API

Bu proje, Go dilinin gücünü ve modern backend mimarisi standartlarını öğrenmek/uygulamak amacıyla geliştirilmiş bir RESTful API'dir. Araçların garaja eklenmesi, güncellenmesi, silinmesi ve bu araçlara modifikasyon parçalarının takılması (One-to-Many ilişkisi) süreçlerini yönetir.

## 🛠️ Kullanılan Teknolojiler

*   **Dil:** [Go (Golang)](https://go.dev/)
*   **Web Framework:** [Fiber v3](https://docs.gofiber.io/)
*   **ORM:** [GORM](https://gorm.io/)
*   **Veritabanı:** SQLite
*   **Doğrulama (Validation):** Go-Playground Validator (v10)
*   **Çevre Değişkenleri:** Godotenv

## 🏗️ Mimari ve Tasarım Desenleri

Proje, kurumsal ölçekte büyüyecek şekilde tasarlanmış "Separation of Concerns" (Sorumlulukların Ayrılığı) prensibine uygun bir klasör yapısına sahiptir:

*   **Models:** Veritabanı şablonları ve tablolar arası ilişkiler (GORM Auto-Migration).
*   **Controllers:** İş mantığının (CRUD işlemleri) yürütüldüğü katman.
*   **DTO (Data Transfer Object):** İstemciden gelen verilerin filtrelenmesi ve doğrulanması (Validation).
*   **Routes:** Route Grouping yapılarak rotaların yönetilebilir hale getirilmesi.
*   **Middleware:** Fiber Logger ile gelen HTTP isteklerinin terminalden izlenmesi.

## 📁 Proje Yapısı

    ├── controllers/
    │   └── car_controller.go
    ├── database/
    │   └── database.go
    ├── dto/
    │   └── car_dto.go
    ├── models/
    │   └── car.go
    ├── routes/
    │   └── routes.go
    ├── .env.example
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    └── main.go

## 🚀 Kurulum ve Çalıştırma

Projeyi kendi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyin:

1. Projeyi klonlayın:
   git clone https://github.com/KULLANICI_ADIN/tuning-garage-api.git
   cd tuning-garage-api

2. Gerekli kütüphaneleri indirin:
   go mod download

3. Çevre değişkenlerini ayarlayın:
   * .env.example dosyasının adını .env olarak değiştirin.
   * Gerekirse içindeki PORT veya DB_NAME değerlerini kendinize göre güncelleyin.

4. Uygulamayı başlatın:
   go run main.go
   
   *(Not: Uygulama ilk çalıştığında SQLite veritabanı dosyasını ve tabloları otomatik olarak oluşturacaktır.)*

## 📡 API Uç Noktaları (Endpoints)

| Metod  | Uç Nokta (Endpoint)       | Açıklama                                       |
| :---   | :---                      | :---                                           |
| GET    | `/api/cars`               | Garajdaki tüm araçları ve modifikasyonları getirir. |
| POST   | `/api/cars`               | Garaja yeni bir araç ekler (Validator destekli). |
| PUT    | `/api/cars/:id`           | Garajdaki mevcut bir aracın marka/modelini günceller. |
| DELETE | `/api/cars/:id`           | Aracı garajdan siler (GORM Soft Delete uygular). |
| POST   | `/api/cars/:id/mods`      | Belirtilen araca yeni bir modifikasyon parçası takar. |

---
*Bu proje, backend mühendisliği, veri yapıları ve temiz mimari pratikleri üzerine çalışırken geliştirilmiştir.*

# Courier / Tracking API 🚚

RESTful API untuk sistem manajemen pengiriman dan pelacakan ekspedisi tingkat enterprise.
Dibangun menggunakan **Golang**, **Fiber**, **MySQL**, dan menerapkan konsep **Clean Architecture** serta dokumentasi API menggunakan **Swagger**.

---

# 🚀 Fitur Utama

* 🔐 **Authentication & Authorization**
  JWT Authentication untuk Admin dan Kurir.

* 👥 **Courier Management**
  Entitas kurir terpisah dari user login, lengkap dengan profil, kendaraan, dan status kerja.

* 💰 **Dynamic Zone Pricing**
  Tarif pengiriman otomatis berdasarkan rute dan tipe layanan.

* 📦 **Multi-Item Shipment**
  Satu nomor resi dapat memiliki banyak item barang.

* 📄 **PDF Airway Bill**
  Generate label pengiriman dalam format PDF.

* 📍 **Live Tracking System**
  Pelacakan paket lengkap dengan timeline histori.

* 📷 **Proof of Delivery (POD)**
  Upload foto bukti penerimaan paket.

* 📊 **Dashboard Statistics**
  Statistik pengiriman dan ringkasan data.

* 📑 **Swagger Documentation**
  Dokumentasi API interaktif menggunakan Swagger.

---

# 🛠 Tech Stack

| Technology       | Description                  |
| ---------------- | ---------------------------- |
| Golang           | Backend Programming Language |
| Fiber v2         | Web Framework                |
| GORM             | ORM Library                  |
| MySQL            | Database                     |
| JWT              | Authentication Mechanism     |
| Swagger (Swaggo) | API Documentation            |
| FPDF             | PDF Generation               |

---

# 📁 Project Structure

Project ini menggunakan konsep **Clean Architecture**.

```bash
courier-api/
├── config/          # Database & Environment Configuration
├── constants/       # Static Constants & Status Codes
├── controllers/     # HTTP Handler Layer
├── docs/            # Swagger Generated Files
├── middleware/      # JWT Auth & CORS Middleware
├── models/          # Database Models
├── repositories/    # Data Access Layer
├── services/        # Business Logic Layer
├── utils/           # Helper Functions
├── uploads/         # POD Image Storage
├── .env             # Environment Variables
├── main.go          # Application Entry Point
└── go.mod           # Go Dependencies
```

---

# ⚙️ Installation & Setup

## 1. Prerequisites

Pastikan sudah menginstall:

* Go 1.18+
* MySQL Server
* Git

---

## 2. Clone Repository

```bash
git clone https://github.com/username/courier-api.git
cd courier-api
```

---

## 3. Create Environment File

Buat file `.env`:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=courier_db
JWT_SECRET=rahasia_super_aman
PORT=3000
```

---

## 4. Install Dependencies

```bash
go mod tidy
```

---

## 5. Install Swagger CLI

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate Swagger Documentation:

```bash
swag init -g main.go -o ./docs
```

---

## 6. Run Application

```bash
go run main.go
```

Server akan berjalan di:

```bash
http://localhost:3000
```

---

# 🔑 Default Admin Account

```txt
Username : admin
Password : admin123
```

---

# 📖 Swagger Documentation

Akses Swagger UI melalui:

```bash
http://localhost:3000/swagger/index.html
```

---

# 📌 API Specification

## Base URL

```bash
http://localhost:3000/api/v1
```

---

# 🔐 Authentication

| Method | Endpoint  | Description         | Access |
| ------ | --------- | ------------------- | ------ |
| POST   | /login    | Login Admin / Kurir | Public |
| POST   | /register | Register User Baru  | Admin  |

---

# 👥 Master Data (Admin Only)

## Tariffs

| Method | Endpoint | Description          |
| ------ | -------- | -------------------- |
| POST   | /tariffs | Membuat tarif baru   |
| GET    | /tariffs | Melihat daftar tarif |

---

## Couriers

| Method | Endpoint  | Description          |
| ------ | --------- | -------------------- |
| POST   | /couriers | Membuat profil kurir |
| GET    | /couriers | Melihat daftar kurir |

---

# 📦 Shipment Management

| Method | Endpoint                | Description            | Access        |
| ------ | ----------------------- | ---------------------- | ------------- |
| POST   | /shipments              | Membuat shipment baru  | Admin         |
| GET    | /shipments              | List shipment          | Admin / Kurir |
| POST   | /shipments/:resi/assign | Assign kurir           | Admin         |
| PATCH  | /shipments/:resi/status | Update status shipment | Kurir         |
| POST   | /shipments/:resi/pod    | Upload POD             | Kurir         |
| GET    | /shipments/:resi/pdf    | Download PDF Label     | Admin / Kurir |
| GET    | /shipments/stats        | Statistik dashboard    | Admin         |

---

## Example Request - Create Shipment

```json
{
  "sender_name": "PT Maju Jaya",
  "sender_city": "Jakarta",
  "receiver_name": "Budi",
  "receiver_city": "Bandung",
  "service_type": "EXPRESS",
  "items": [
    {
      "item_name": "Laptop",
      "quantity": 1,
      "weight": 2.5
    }
  ]
}
```

---

# 🛵 Courier App

| Method | Endpoint  | Description        | Access |
| ------ | --------- | ------------------ | ------ |
| GET    | /my-tasks | Daftar tugas kurir | Kurir  |

---

# 📍 Public Tracking

| Method | Endpoint     | Description               | Access |
| ------ | ------------ | ------------------------- | ------ |
| GET    | /track/:resi | Tracking paket & timeline | Public |

---

# 🗃️ Database Schema

* **Users** → Data login user.
* **Couriers** → Profil kerja kurir.
* **Tariffs** → Tarif pengiriman.
* **Shipments** → Data utama pengiriman.
* **ShipmentItems** → Detail item barang.
* **TrackingHistories** → Histori perjalanan paket.

---

# 📌 Shipment Status

| Status           | Description            |
| ---------------- | ---------------------- |
| PENDING          | Shipment baru dibuat   |
| COURIER_ASSIGNED | Kurir telah ditugaskan |
| IN_TRANSIT       | Paket sedang dikirim   |
| DELIVERED        | Paket telah diterima   |

---

# 📌 Courier Status

| Status  | Description               |
| ------- | ------------------------- |
| OFFLINE | Kurir tidak aktif         |
| ONLINE  | Kurir siap menerima tugas |

---

# 🚀 Future Improvements

* Docker & Docker Compose
* Unit Testing & Integration Testing
* Redis Caching
* Rate Limiter
* Email / SMS Notification
* Real-time Tracking via WebSocket

---

# A little bit of logic, a little bit of love ❤️

Built using Golang & Fiber.
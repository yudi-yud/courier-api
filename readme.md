# Courier / Tracking API 🚚

RESTful API untuk sistem manajemen pengiriman dan pelacakan (tracking) ekspedisi tingkat lanjut.  
API ini dibangun menggunakan bahasa pemrograman Go (Golang), framework Fiber, database MySQL, serta menerapkan Clean Architecture untuk menjaga struktur kode tetap scalable dan mudah di-maintain.

---

# 🚀 Fitur Utama

- 🔐 Authentication & Authorization menggunakan JWT
- 👥 Manajemen Kurir (Courier Management)
- 💰 Tarif Dinamis berdasarkan kota asal, tujuan, dan tipe layanan
- 📦 Multi-item shipment dalam satu nomor resi
- 📄 Generate PDF Airway Bill otomatis
- 📍 Live tracking shipment timeline
- 📷 Upload POD (Proof of Delivery)
- 📊 Dashboard statistics shipment
- 🔍 Pagination & Search shipment
- 🧱 Clean Architecture (Controller, Service, Repository)

---

# 🛠 Tech Stack

| Technology | Description |
|---|---|
| Golang | Backend Programming Language |
| Fiber v2 | Web Framework |
| GORM | ORM Library |
| MySQL | Database |
| JWT | Authentication Mechanism |
| FPDF | PDF Generation Library |

---

# 📁 Struktur Folder

```bash
courier-api/
├── config/              # Database & Environment Configuration
├── controllers/         # HTTP Handler Layer
├── middleware/          # JWT Middleware & CORS
├── models/              # GORM Models
├── repositories/        # Data Access Layer
├── services/            # Business Logic Layer
├── utils/               # Helper Functions
├── uploads/             # POD Image Storage
├── .env
├── main.go
└── go.mod
```

---

# ⚙️ Instalasi & Setup

## 1. Prasyarat

Pastikan sudah terinstall:

- Go 1.18+
- MySQL Server
- Git

---

## 2. Clone Repository

```bash
git clone https://github.com/username/courier-api.git

cd courier-api
```

---

## 3. Konfigurasi Environment

Buat file `.env`

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=courier_db

JWT_SECRET=your_super_secret_key

PORT=3000
```

---

## 4. Buat Database

```sql
CREATE DATABASE courier_db;
```

---

## 5. Jalankan Aplikasi

```bash
# Install dependencies
go mod tidy

# Run application
go run main.go
```

Server berjalan di:

```bash
http://localhost:3000
```

Aplikasi otomatis:
- Melakukan migrasi database
- Membuat user admin default

---

# 👤 Default Admin

```text
Username : admin
Password : admin123
```

---

# 📖 API Specification

## Base URL

```bash
http://localhost:3000/api/v1
```

---

# 🔐 Authentication

## Login

Mendapatkan JWT access token.

### Endpoint

```http
POST /login
```

### Request Body

```json
{
  "username": "admin",
  "password": "admin123"
}
```

### Response

```json
{
  "success": true,
  "code": 200,
  "message": "Login success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

---

## Register User (Admin Only)

Membuat user login baru untuk Admin atau Courier.

### Endpoint

```http
POST /register
```

### Header

```http
Authorization: Bearer <token>
```

### Request Body

```json
{
  "username": "courier_baru",
  "password": "password123",
  "role": "courier"
}
```

---

# 👥 Courier Management

## Create Courier Profile

Membuat profil courier berdasarkan user login.

### Endpoint

```http
POST /couriers
```

### Header

```http
Authorization: Bearer <token>
```

### Request Body

```json
{
  "name": "Jono Supir",
  "phone": "08123456789",
  "vehicle_plate": "B 1234 XX",
  "vehicle_type": "Motor",
  "user_id": 2
}
```

---

## Get All Couriers

### Endpoint

```http
GET /couriers
```

### Header

```http
Authorization: Bearer <token>
```

---

# 💰 Tariff Management

## Create Tariff

Membuat tarif ongkir berdasarkan kota asal, tujuan, dan tipe layanan.

### Endpoint

```http
POST /tariffs
```

### Request Body

```json
{
  "origin_city": "Jakarta",
  "destination_city": "Bandung",
  "service_type": "EXPRESS",
  "price_per_kg": 20000,
  "etd": "1-2 Days"
}
```

---

## Get All Tariffs

### Endpoint

```http
GET /tariffs
```

---

# 📦 Shipment Management

## Create Shipment (Admin Only)

Membuat shipment baru dengan:
- Auto generate resi
- Dynamic pricing
- Multi-item shipment

### Endpoint

```http
POST /shipments
```

### Header

```http
Authorization: Bearer <token>
```

### Request Body

```json
{
  "sender_name": "PT Maju Jaya",
  "sender_phone": "021111222",
  "sender_address": "Jl. Industri No.5",
  "sender_city": "Jakarta",

  "receiver_name": "Budi Penerima",
  "receiver_phone": "0812999888",
  "receiver_address": "Jl. Dago Atas No.10",
  "receiver_city": "Bandung",

  "service_type": "EXPRESS",

  "items": [
    {
      "item_name": "Laptop",
      "quantity": 1,
      "weight": 2.5
    },
    {
      "item_name": "Kardus Besar",
      "quantity": 2,
      "weight": 5.0
    }
  ]
}
```

### Response

```json
{
  "success": true,
  "code": 201,
  "message": "Shipment created successfully",
  "data": {
    "resi_number": "TRK123456",
    "total_weight": 7.5,
    "price": 150000,
    "service_type": "EXPRESS",
    "etd": "1-2 Days",
    "status": "PENDING"
  }
}
```

---

## Get All Shipments

Mendukung pagination dan search.

### Endpoint

```http
GET /shipments?page=1&limit=10&search=Budi
```

### Header

```http
Authorization: Bearer <token>
```

---

## Generate PDF Airway Bill

Generate label pengiriman dalam format PDF.

### Endpoint

```http
GET /shipments/:resi/pdf
```

### Header

```http
Authorization: Bearer <token>
```

### Response

```http
Content-Type: application/pdf
```

---

## Assign Courier (Admin Only)

Menugaskan courier ke shipment.

### Endpoint

```http
POST /shipments/:resi/assign
```

### Request Body

```json
{
  "courier_id": 1
}
```

---

## Update Shipment Status (Courier)

### Endpoint

```http
PATCH /shipments/:resi/status
```

### Header

```http
Authorization: Bearer <token>
```

### Request Body

```json
{
  "status": "IN_TRANSIT",
  "location": "Tol KM 50",
  "note": "Menuju kota tujuan"
}
```

---

## Upload POD (Proof of Delivery)

Upload bukti foto penerimaan paket.

### Endpoint

```http
POST /shipments/:resi/pod
```

### Header

```http
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

### Form Data

| Key | Type |
|---|---|
| pod_image | File |

---

# 📍 Public Tracking

## Track Shipment by Resi

Endpoint publik untuk tracking shipment tanpa login.

### Endpoint

```http
GET /track/:resi
```

### Response

```json
{
  "success": true,
  "code": 200,
  "message": "Shipment found",
  "data": {
    "resi_number": "TRK123456",
    "status": "DELIVERED",

    "courier": {
      "name": "Jono Supir",
      "phone": "08123456789"
    },

    "items": [
      {
        "item_name": "Laptop",
        "quantity": 1,
        "weight": 2.5
      }
    ],

    "histories": [
      {
        "status": "PENDING",
        "timestamp": "2026-05-21T10:00:00Z",
        "note": "Shipment created"
      },
      {
        "status": "DELIVERED",
        "timestamp": "2026-05-21T15:00:00Z",
        "note": "Package delivered"
      }
    ]
  }
}
```

---

# 🗃️ Database Schema

## Tables

1. Users
2. Couriers
3. Tariffs
4. Shipments
5. ShipmentItems
6. TrackingHistories

---

# 📌 Shipment Status

| Status | Description |
|---|---|
| PENDING | Shipment baru dibuat |
| COURIER_ASSIGNED | Courier sudah ditugaskan |
| PICKED_UP | Paket diambil courier |
| IN_TRANSIT | Paket dalam perjalanan |
| ARRIVED_AT_HUB | Paket tiba di hub |
| OUT_FOR_DELIVERY | Paket dikirim ke penerima |
| DELIVERED | Paket diterima |
| CANCELLED | Shipment dibatalkan |

---

# 📌 HTTP Status Code Reference

| Code | Description |
|---|---|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 500 | Internal Server Error |

---

# 📄 Standard JSON Response

## Success Response

```json
{
  "success": true,
  "code": 200,
  "message": "Success",
  "data": {}
}
```

## Error Response

```json
{
  "success": false,
  "code": 400,
  "message": "Validation error",
  "errors": {
    "field": [
      "field is required"
    ]
  }
}
```

---

# 🧪 API Testing

Disarankan menggunakan:

- Postman
- Insomnia

## Langkah Testing

1. Login terlebih dahulu
2. Copy JWT token
3. Tambahkan token ke header:

```http
Authorization: Bearer <token>
```

# A little bit of logic, a little bit of love ❤️

Built using Golang & Fiber.
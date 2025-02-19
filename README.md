# Go-Gin-Mongo-App

## 📌 Project Overview
This is a **feature-based** Go web server using **Gin** and **MongoDB**, following best practices with **repository-service-controller** structure. It also supports **API versioning** and **authentication middleware**.

### ✅ Features
- 🌍 **RESTful API using Gin**
- 🔐 **Authentication (Register/Login)**
- 🔒 **JWT-based Auth Middleware**
- 🗄 **MongoDB ORM-like Repository Pattern**
- 📌 **Feature-based Modular Structure**
- 🏥 **Health Check Endpoint**
- 🛡 **User Info API (Protected)**

---

## 📂 Folder Structure

```
go-gin-mongo-app/
│── main.go
│── go.mod
│── go.sum
│── .env                 # Environment variables
│── config/              # Database config
│── routes/              # API route registration
│── middleware/          # Authentication middleware
│── common/models/       # Shared database models
│── domain/
│   ├── auth/v1/         # Authentication API (v1)
│   ├── userInfo/v1/     # User Info API (v1, protected)
│   ├── health/v1/       # Health Check API (v1, public)
└── utils/               # Utility functions (hashing, etc.)
```

---

## 🚀 Setup & Installation

### 1️⃣ **Clone Repository**
```sh
git clone https://github.com/kenchimaru/go-gin-features-based.git
cd go-gin-features-based
```

### 2️⃣ **Set Up Environment Variables (.env)**
Create a `.env` file in the project root and add the following:
```
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your-secret-key
```

### 3️⃣ **Install Dependencies**
```sh
go mod tidy
```

### 4️⃣ **Run the Application**
```sh
go run main.go
```

The server will start on **`http://localhost:8080`**.

---

## 📌 API Endpoints

### 🟢 **Public API**
| Method | Endpoint               | Description        |
|--------|------------------------|--------------------|
| `GET`  | `/api/v1/healthcheck`  | Health Check API  |

### 🔒 **Authentication API**
| Method  | Endpoint               | Description        |
|---------|------------------------|--------------------|
| `POST`  | `/api/v1/auth/register` | Register User     |
| `POST`  | `/api/v1/auth/login`    | Login User        |

### 🔐 **Protected User Info API (Requires Auth)**
| Method  | Endpoint            | Description     |
|---------|---------------------|---------------|
| `GET`   | `/api/v1/user/info` | Get User Info |

---

## 🛠 **Technologies Used**
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [JWT Authentication](https://pkg.go.dev/github.com/dgrijalva/jwt-go)
- [Bcrypt for Password Hashing](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

## 🎯 Contributing
Feel free to contribute by creating PRs! If you find a bug, open an issue.

---

## 📜 License
This project is licensed under the MIT License.

# 🛒 Store API (Golang + MongoDB)

A modern and secure RESTful API built with **Golang**, **Gin** framework, and **MongoDB** for managing user registration, authentication, and protected user routes.

---

## ✨ Features

- 🔐 Secure user registration and login with bcrypt password hashing
- 🪪 JWT-based authentication with middleware protection
- 📦 MongoDB integration for persistent storage
- 🔍 Protected `/me` route to fetch authenticated user info
- ⚙️ Clean project structure following best practices

---

## 🚀 Endpoints

### 🔸 Register

`POST /api/v1/register`

**Body:**

```json
{
  "name": "Aliakbar Zohour",
  "email": "your@email.com",
  "password": "yourpassword"
}
```

### 🔸 Login

`POST /api/v1/login`

**Body:**

```json
{
  "email": "your@email.com",
  "password": "yourpassword"
}
```

**Returns:**

```json
{
  "token": "<JWT_TOKEN>"
}
```

### 🔒 Get Authenticated User

`GET /api/v1/me`

**Headers:**

```
Authorization: Bearer <JWT_TOKEN>
```

**Returns:**

```json
{
  "name": "Aliakbar Zohour",
  "email": "your@email.com",
  "role": "user"
}
```

---

## 🛠️ Technologies

- Golang
- Gin Gonic
- MongoDB
- JWT
- Bcrypt

---

## 🧩 Project Structure

```
.
├── cmd
│   └── main.go             # Entry point
├── config
│   └── db.go               # MongoDB connection setup
├── internal
│   ├── handlers            # Route handlers (register, login, me)
│   ├── middleware          # Auth middleware
│   ├── models              # User model definition
│   ├── repository          # DB interaction functions
│   └── routes              # Route grouping
├── pkg
│   └── utils.go            # JWT, Hashing helpers
└── go.mod / go.sum
```

---

## 📦 Installation

```bash
git clone https://github.com/aliakbar-zohour/store-api.git
cd store-api
go mod tidy
go run cmd/main.go
```

---

## 🧪 Testing with cURL

```bash
# Register
curl -X POST http://localhost:8080/api/v1/register -H "Content-Type: application/json" -d '{"name":"Aliakbar","email":"az@gmail.com","password":"admin"}'

# Login
curl -X POST http://localhost:8080/api/v1/login -H "Content-Type: application/json" -d '{"email":"az@gmail.com","password":"admin"}'

# Access /me
curl -X GET http://localhost:8080/api/v1/me -H "Authorization: Bearer <your-token>"
```

---

## 🧑‍💻 Author

**Aliakbar Zohour**
Built with ❤️ using Go and MongoDB

---

## 📄 License

MIT License

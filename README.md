# gin-app-auth

`gin-app-auth` is a lightweight, high-performance authentication microservice built using the [Gin Web Framework](https://github.com/gin-gonic/gin) in Go. It provides RESTful APIs for user registration, login, token-based authentication (JWT), and user management, making it suitable for modern microservices or monolithic applications that need an authentication layer.

---

## ✨ Features

- 🔐 JWT-based authentication (access & refresh tokens)
- 👤 User registration and login
- 🧼 Password hashing using bcrypt
- 🛡️ Role-based access control (optional)
- ✅ Input validation
- 🧪 Easy to test and extend
- ⚙️ Configurable via environment variables
- 📦 Follows SOLID and Clean Architecture principles (optional)

---

## 📁 Project Structure
```bash
└───GreatThanosApp
    │   .env
    │   dockerfile
    │   go.mod
    │   go.sum
    │   mac.env
    │   windows.env
    │   __debug_bin150744087.exe
    │
    ├───.github
    │   └───workflows
    │           go.yml
    │
    ├───.vscode
    │       launch.json
    │
    ├───api
    │   ├───handler
    │   │       auth_handler.go
    │   │       user_handler.go
    │   │
    │   ├───middleware
    │   │       auth_middleware.go
    │   │
    │   └───router
    │           router.go
    │
    ├───cmd
    │       main.go
    │
    ├───config
    │       config.go
    │       database.go
    │
    ├───docs
    │       docs.go
    │       swagger.json
    │       swagger.yaml
    │
    ├───internal
    │   ├───dto
    │   │       auth_dto.go
    │   │       user_dto.go
    │   │
    │   ├───repository
    │   │       auth_repository.go
    │   │       user_repository.go
    │   │
    │   ├───service
    │   │       auth_service.go
    │   │       user_service.go
    │   │
    │   └───usecase
    │           auth_usecase.go
    │           user_usecase.go
    │
    ├───k8s
    │       deployment.yaml
    │       service.yaml
    │
    ├───models
    │       user.go
    │       userLogin.go
    │
    ├───pkg
    │   │   response.go
    │   │
    │   ├───external
    │   └───logger
    └───utils
            jwt.go
            password.go
            phone.go
            username.go
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- (Optional) PostgreSQL or MongoDB or any supported database
- `go mod` enabled

### Installation

```bash
git clone https://github.com/your-username/gin-app-auth.git
cd gin-app-auth
go mod tidy

# Go API with MongoDB Atlas

A simple REST API built with Go (Gin) and MongoDB Atlas.

## Features

- ✅ Create, Read, Update, Delete (CRUD) operations
- ✅ MongoDB Atlas integration
- ✅ CORS enabled
- ✅ Clean code structure

## Prerequisites

- Go 1.18+
- MongoDB Atlas account (free tier)
- Git

## Setup

### 1. Clone Repository
```bash
git clone https://github.com/YOUR_USERNAME/go-api-mongodb.git
cd go-api-mongodb
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Create .env File
Create a `.env` file in root directory:
MONGODB_URI=mongodb+srv://admin:YOUR_PASSWORD@cluster0.xxxxx.mongodb.net/?retryWrites=true&w=majority
DATABASE_NAME=go_api_db
PORT=8000

### 4. Run the API
```bash
go run main.go
```

Server starts at: `http://localhost:8000`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users` | Create new user |
| GET | `/api/users` | Get all users |
| GET | `/api/users/:id` | Get user by ID |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Delete user |

## Example Requests

### Create User
```bash
curl -X POST http://localhost:8000/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":25}'
```

### Get All Users
```bash
curl http://localhost:8000/api/users
```

### Get User by ID
```bash
curl http://localhost:8000/api/users/64a1b2c3d4e5f6g7h8i9j0k1
```

### Update User
```bash
curl -X PUT http://localhost:8000/api/users/64a1b2c3d4e5f6g7h8i9j0k1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com","age":26}'
```

### Delete User
```bash
curl -X DELETE http://localhost:8000/api/users/64a1b2c3d4e5f6g7h8i9j0k1
```

## Project Structure
go-api-mongodb/
├── main.go                 # Entry point
├── config/                 # Configuration
│   └── config.go
├── database/               # MongoDB connection
│   └── database.go
├── handlers/               # API handlers
│   └── user_handler.go
├── models/                 # Data models
│   └── user.go
├── middleware/             # Custom middleware
│   └── cors.go
├── go.mod                  # Dependencies
├── .env                    # Environment variables
├── .gitignore             # Git ignore rules
└── README.md              # This file

## Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **MongoDB** - Database
- **MongoDB Atlas** - Managed MongoDB service

## Author

Kristine Joy Martinez (@kjvmartinez)

## License

MIT License
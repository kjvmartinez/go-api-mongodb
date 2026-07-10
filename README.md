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

### Error Responses:

**Missing Authorization Header:**
```json
{
  "error": "missing authorization header"
}
```

**Invalid Token:**
```json
{
  "error": "invalid token"
}
```

**Unauthorized:**
```json
{
  "error": "invalid email or password"
}
```


## Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **MongoDB** - Database
- **MongoDB Atlas** - Managed MongoDB service

## Author

Kristine Joy Martinez (@kjvmartinez)

## License

MIT License
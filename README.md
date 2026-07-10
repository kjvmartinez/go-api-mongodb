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

### Authentication Endpoints (No Auth Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/register` | Register new user |
| POST | `/api/auth/login` | Login user |

### User Endpoints (Auth Required)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|----------------|
| POST | `/api/users` | Create new user | ✅ Yes |
| GET | `/api/users` | Get all users | ✅ Yes |
| GET | `/api/users/:id` | Get user by ID | ✅ Yes |
| PUT | `/api/users/:id` | Update user | ✅ Yes |
| DELETE | `/api/users/:id` | Delete user | ✅ Yes |

## Example Requests

### Register User

```bash
curl -X POST http://localhost:8000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name":"John Doe",
    "email":"john@example.com",
    "password":"password123",
    "age":25
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "507f1f77bcf86cd799439011",
    "name": "John Doe",
    "email": "john@example.com",
    "age": 25
  }
}
```

### Login User

```bash
curl -X POST http://localhost:8000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email":"john@example.com",
    "password":"password123"
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "507f1f77bcf86cd799439011",
    "name": "John Doe",
    "email": "john@example.com",
    "age": 25
  }
}
```

### Create User (Protected)

```bash
curl -X POST http://localhost:8000/api/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"name":"Jane Doe","email":"jane@example.com","age":26}'
```

### Get All Users (Protected)

```bash
curl -X GET http://localhost:8000/api/users \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Get User by ID (Protected)

```bash
curl -X GET http://localhost:8000/api/users/507f1f77bcf86cd799439011 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Update User (Protected)

```bash
curl -X PUT http://localhost:8000/api/users/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"name":"Jane Smith","email":"jane.smith@example.com","age":27}'
```

### Delete User (Protected)

```bash
curl -X DELETE http://localhost:8000/api/users/507f1f77bcf86cd799439011 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Authentication

This API uses **JWT (JSON Web Token)** for authentication.

### How to Authenticate:

1. **Register or Login** to get a token
2. **Add token to requests** using Authorization header:
  Authorization: Bearer YOUR_TOKEN_HERE
3. **Token is valid for 24 hours**

### Headers for Protected Routes:
  Content-Type: application/json
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
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
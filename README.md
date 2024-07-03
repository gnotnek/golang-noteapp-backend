# Golang Note Management Backend Service

## 1. Project Overview

This project is a backend service for managing notes with user accounts, built using Golang. It provides functionality for user authentication, note creation, retrieval, updating, and deletion. The service uses PostgreSQL as the database and implements JWT for authentication.

## 2. Installation

### Prerequisites:
- Go (version 1.x or higher)
- PostgreSQL (version x.x or higher)

### Steps:

1. Clone the repository:
   ```
   git clone https://github.com/gnotnek/golang-noteapp-backend.git
   cd golang-noteapp-backend
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up the environment:
   - Copy the `.env.example` file to `.env`
   - Update the `.env` file with your PostgreSQL credentials and JWT key:
     ```
     DB_HOST=your_host
     DB_PORT=your_port
     DB_USER=your_username
     DB_NAME=your_database_name
     DB_PASSWORD=your_password
     DB_SSLMODE=disable
     JWT_KEY=your_secret_key
     ```

4. Run the project:
   ```
   go run main.go
   ```

The server should now be running on `http://localhost:8080`.

## 3. Usage

The API provides endpoints for user registration, login, and note management. All note-related endpoints require authentication.

### Authentication

To access protected endpoints, include the JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## 4. API Reference

### User Endpoints

#### Register a new user
- **POST** `/api/register`
- Body:
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```

#### Login
- **POST** `/api/login`
- Body:
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- Response: JWT token

### Note Endpoints

All note endpoints require authentication.

#### Create a new note
- **POST** `/api/notes`
- Body:
  ```json
  {
    "title": "Note Title",
    "body": "Note Content",
    "userId": 1
  }
  ```

#### Get all notes
- **GET** `/api/notes`

#### Get a specific note
- **GET** `/api/notes/:id`

#### Update a note
- **PUT** `/api/notes/:id`
- Body:
  ```json
  {
    "title": "Updated Title",
    "body": "Updated Content"
  }
  ```

#### Delete a note
- **DELETE** `/api/notes/:id`

## 5. Architecture

The project follows a typical MVC (Model-View-Controller) architecture with some additional layers:

### Main Components:

1. **Models**: Define the data structures (User and Notes).
2. **Controllers**: Handle incoming HTTP requests and manage the application logic.
3. **Services**: Implement the business logic and database operations.
4. **Middleware**: Implement authentication checks.
5. **Routes**: Define the API endpoints and connect them to controllers.
6. **Database**: Manages the connection to PostgreSQL using GORM ORM.
7. **Auth**: Handles JWT token creation and verification.

### Key Files and Their Roles:

- `main.go`: Entry point of the application. Initializes the database and sets up the router.
- `routes/route.go`: Defines all API routes and applies middleware.
- `services/postgresql/database.go`: Manages database connection and initialization.
- `services/postgresql/user_service.go`: Implements user-related database operations.
- `services/postgresql/notes_service.go`: Implements note-related database operations.
- `controllers/user_controller.go`: Handles user registration and login.
- `controllers/note_controller.go`: Manages CRUD operations for notes.
- `middleware/auth.go`: Implements JWT authentication middleware.
- `auth/jwt.go`: Provides functions for creating and verifying JWT tokens.
- `models/user_model.go` and `models/note_model.go`: Define the data structures for users and notes.

The application uses the Gin web framework for handling HTTP requests and GORM as the ORM for database operations. JWT is used for stateless authentication.

## 6. Database Schema

The application uses two main tables:

1. **Users**:
   - ID (UUID, primary key)
   - Username (string, unique)
   - Password (string, hashed)

2. **Notes**:
   - ID (UUID, primary key)
   - Title (string)
   - Body (string)
   - UserID (uint, foreign key to Users)

The schema is automatically managed by GORM's AutoMigrate feature.
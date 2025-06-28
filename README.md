# Credible Mandela API

A Go-based backend service for a music application that provides podcast management, user authentication, and search functionality using MongoDB.

## Features

- User authentication and authorization
- Community notes management (CRUD operations)
- Like/Unlike functionality
- User-specific note retrieval
- Advertisement management system
- MongoDB integration
- Docker support

## Prerequisites

- Go 1.22.3 or higher
- MongoDB
- Docker and Docker Compose

## Installation

1.  Install dependencies:
    ```
    go mod download
    ```
2.  Set up environment variables: Create a .env file in the root directory with the following variables:
    ```
    MONGODB_URI=mongodb://localhost:27017
    PORT=8080
    JWT_SECRET=your-secret-key
    ```

## Project Structure

```
credible-mandela-api/
├── config/         # Configuration management
├── controllers/    # Request handlers
├── middlewares/    # Custom middleware
├── models/         # Data models
├── routers/        # Route definitions
├── services/       # Business logic
├── utils/          # Helper functions
├── docker-compose.yaml
├── go.mod
├── go.sum
└── main.go
```

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user

### Community Notes

- `GET /api/community-notes` - Get all community notes
- `GET /api/community-notes/:id` - Get specific community note
- `POST /api/community-notes` - Create new community note
- `DELETE /api/community-notes/:id` - Delete community note
- `POST /api/community-notes/like/:id` - Like a community note
- `POST /api/community-notes/unlike/:id` - Unlike a community note
- `GET /api/community-notes/user/:username` - Get user's community notes
- `GET /api/community-notes/user/me` - Get current user's notes

### Advertisements

- `GET /api/ads` - Get all advertisements
- `GET /api/ads/:id` - Get specific advertisement by ID
- `GET /api/ads/user/:address` - Get all advertisements by user address
- `GET /api/ads/user/me` - Get current user's advertisements
- `POST /api/ads` - Publish new advertisement
- `PUT /api/ads/:id` - Update advertisement
- `DELETE /api/ads/:id` - Delete advertisement

## Running the Application

1.  Start MongoDB:
    ```
    docker-compose up -d
    ```
2.  Run the application:
    `    go run main.go
   `
    The server will start on http://localhost:8080 by default.

## Example API Usage

### User Registration

```
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

### User Login

```
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

### Create Community Note

```
curl -X POST http://localhost:8080/api/community-notes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{"title": "My Note", "content": "A great note"}'
```

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.

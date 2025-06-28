# Credible Mandela API

A robust Go-based REST API for managing community notes and user interactions, built with a focus on security and performance.

🌐 **Website**: [crediblemandela.xyz](https://tek-zeki-sensin.vercel.app/community-notes)

## 🌟 Key Features & Technologies

### Core Functionality

- **User Authentication**: Secure user registration and login with JWT (access and refresh tokens).
- **Community Notes**: Full CRUD (Create, Read, Update, Delete) operations for community-driven notes.
- **Engagement**: Like and Unlike functionality for notes.
- **Advertisement System**: Manage and display advertisements within the platform.

### Technical Stack

- **Language**: **Go (Golang)**
- **Framework**: **Gin Web Framework** for high-performance routing and middleware.
- **Database**: **MongoDB** with the official Go driver for flexible, scalable data storage.
- **Containerization**: **Docker** and **Docker Compose** for consistent development and deployment environments.
- **Configuration**: **Viper** for managing application configuration from environment variables and files.
- **Authentication**: **JWT (JSON Web Tokens)** with **RSA (RS256)** signing for secure, stateless authentication.

### Security

- **Asymmetric-Key Cryptography**: Uses **RSA keys** for signing and verifying JWTs, ensuring token integrity.
- **Password Hashing**: Employs **bcrypt** for securely hashing and storing user passwords.
- **CORS**: Configured Cross-Origin Resource Sharing (CORS) to control access from different domains.
- **Input Validation**: Validates incoming request data to prevent common vulnerabilities.

## 🏗️ Project Structure

The project follows a clean, modular architecture to separate concerns and improve maintainability.

```
credible-mandela-api/
├── config/         # Configuration management (Viper)
├── controllers/    # HTTP request handlers (Gin)
├── middlewares/    # Custom middleware (e.g., authentication)
├── models/         # Data structures and database models
├── routers/        # API route definitions
├── services/       # Core business logic
├── utils/          # Utility functions (crypto, database, etc.)
├── docker-compose.yaml # Docker Compose configuration
├── go.mod          # Go module dependencies
└── main.go         # Application entry point
```

## Prerequisites

- Go 1.22.3 or higher
- Docker and Docker Compose
- MongoDB

## 🚀 Getting Started

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/credible-mandela-api.git
    cd credible-mandela-api
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Set up environment variables:**
    Create a `.env` file in the root directory and add the following:

    ```env
    MONGO_URI=mongodb://localhost:27017
    PORT=8080

    ACCESS_TOKEN_PRIVATE_KEY="your_base64_encoded_private_key"
    ACCESS_TOKEN_PUBLIC_KEY="your_base64_encoded_public_key"
    REFRESH_TOKEN_PRIVATE_KEY="your_base64_encoded_private_key"
    REFRESH_TOKEN_PUBLIC_KEY="your_base64_encoded_public_key"
    ```

### Running the Application

1.  **Start the database:**

    ```bash
    docker-compose up -d
    ```

2.  **Run the Go server:**
    ```bash
    go run main.go
    ```
    The API will be available at `http://localhost:8080`.

## 🔄 API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user.
- `POST /api/auth/login` - Log in and receive JWTs.
- `POST /api/auth/refresh` - Refresh an expired access token.

### Community Notes

- `GET /api/community-notes` - Get all notes.
- `POST /api/community-notes` - Create a new note (requires auth).
- `DELETE /api/community-notes/:id` - Delete a note (requires auth).
- `POST /api/community-notes/like/:id` - Like a note (requires auth).

### Advertisements

- `GET /api/ads` - Get all advertisements.
- `POST /api/ads` - Create a new advertisement (requires auth).

## 💡 Example API Usage

### User Registration

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "address": "0xYourEthereumAddress"
  }'
```

### User Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### Create a Community Note (Authentication Required)

First, log in to get an access token. Then, use the token in the Authorization header.

```bash
# Replace <your-jwt-token> with the actual access token from the login response
curl -X POST http://localhost:8080/api/community-notes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "title": "My First Note",
    "content": "This is a sample community note.",
    "url": "http://example.com"
  }'
```

# Credible Mandela API

A Go-based backend service for a music application that provides podcast management, user authentication, and search functionality using MongoDB.

🌐 **Website**: [crediblemandela.xyz](https://www.crediblemandela.xyz/)

## 🌟 Features

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

## 🏗️ Project Structure

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

## 🔒 Security Features

- JWT-based authentication with ECDSA signing
- Elliptic Curve Digital Signature Algorithm (ECDSA) for cryptographic operations
- CORS configuration
- Request validation
- Secure environment variable management
- Public/private key pair authentication using ECDSA P-256 curve

## 📦 Dependencies

Key dependencies include:

- Gin Web Framework
- MongoDB Go Driver
- JWT Go
- Viper
- CORS middleware

For a complete list, see `go.mod`.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👥 Authors

- Your Name (@yourusername)

## 🙏 Acknowledgments

- Thanks to all contributors
- Inspired by community note-sharing platforms

This README provides:

- Clear installation instructions
- API endpoint documentation
- Project structure overview
- Security features
- Contributing guidelines
- License information

You may want to customize:

1. The repository URL
2. Author information
3. License details
4. Any specific deployment instructions for your environment
5. Add any additional features or requirements specific to your implementation

Would you like me to modify any section or add more specific details to any part?

# healthmate-backend

A backend service that leverages the following technologies:

- **Gin**: Fast, minimalist web framework for Go, used for building RESTful APIs.
- **GORM**: ORM library for Go, used for database interactions.
- **Redis**: In-memory key-value store, used for caching and session management.
- **Logrus**: Structured logger for Go, supports leveled logging and easy integration.
- **JWT**: JSON Web Tokens for stateless authentication and secure API access.

## Features

- RESTful APIs using Gin framework.
- Database operations via GORM.
- Caching and session management using Redis.
- Structured logging with Logrus.
- JWT-based authentication and authorization.

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/tranthanhsang2k3/healthmate-backend.git
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Configure environment variables**  
   Create a `.env` file with necessary configuration for database, Redis, and JWT secret.

4. **Run the application**
   ```bash
   go run main.go
   ```

## Technologies

| Name    | Usage                                      |
|---------|--------------------------------------------|
| Gin     | HTTP server and routing                    |
| GORM    | Database ORM (MySQL/PostgreSQL/SQLite etc) |
| Redis   | Caching, sessions                          |
| Logrus  | Logging                                    |
| JWT     | Authentication, authorization              |

## License

MIT

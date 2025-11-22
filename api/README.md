# AleGYM API Service

The AleGYM API is the backend powerhouse for the AleGYM application. Built with performance and scalability in mind, it provides a secure and efficient RESTful interface for the frontend client.

## Technology Stack

*   **Language:** Go (Golang) v1.25+
*   **Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin) - Fast, low-allocation HTTP web framework.
*   **Database Driver:** [MongoDB Go Driver](https://go.mongodb.org/mongo-driver) - Official driver for MongoDB.
*   **Authentication:** JWT (JSON Web Tokens) with Argon2 hashing for password security.

## Features

*   **User Authentication:** Secure signup, login, and session management using JWTs.
*   **Goal Management:** CRUD operations for user goals and sub-goals.
*   **Progress Tracking:** Endpoints to log and retrieve user progress metrics.
*   **Notes System:** Create, update, and delete personal workout notes.
*   **Event Scheduling:** Manage calendar events and workout sessions.
*   **CORS Support:** Configurable CORS middleware for secure cross-origin requests.

## Configuration

The application relies on environment variables for configuration. Ensure these are set in your environment or a `.env` file in the root directory.

| Variable | Description | Example |
| :--- | :--- | :--- |
| `DATABASE_URL` | Connection string for MongoDB | `mongodb://localhost:27017` |
| `API_CORS_ORIGIN` | Allowed origin for CORS | `http://localhost:5173` |
| `API_PORT` | Port for the API server | `8081` |

## Installation & Running

### Prerequisites
- Go 1.25 or higher installed.
- A running MongoDB instance.

### Steps

1.  **Navigate to the API directory:**
    ```bash
    cd api
    ```

2.  **Install Dependencies:**
    ```bash
    go mod download
    ```

3.  **Run the Server:**
    ```bash
    go run main.go
    ```
    The server will start on port `8081` (or the port specified in your config).

## Project Structure

*   `main.go`: Entry point of the application. Initializes DB connection and sets up the router.
*   `Database/`: Contains database connection logic and data structures (models).
*   `Routes/`: Defines the HTTP route handlers and middleware.
*   `Environment/`: Handles environment variable loading and configuration.

## API Endpoints Overview

*   `POST /session`: Login/Create session.
*   `GET /user`: Get current user profile.
*   `GET /goals`: Retrieve user goals.
*   `POST /goals`: Create a new goal.
*   `GET /notes`: Retrieve user notes.
*   `GET /progress`: Retrieve progress logs.

## License

This project is licensed under the MIT License.

Powered by [yxl-prz/YSvelGoK](https://github.com/yxl-prz/YSvelGoK).

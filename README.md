# AleGYM - GYM Activity Tracker

AleGYM is a comprehensive web application designed to help users track their gym activities, set and monitor goals, and maintain a consistent workout routine. It combines a robust backend service with a modern, responsive frontend to provide a seamless user experience for fitness enthusiasts.

## Project Overview

The goal of this application is to assist users in getting started and staying on track with their gym journey. It goes beyond simple rep counting to include features for:
- **Goal Tracking:** Define main goals and break them down into manageable sub-goals.
- **Progress Monitoring:** Track weight, body metrics, and other progress indicators.
- **Notes & Journaling:** Keep track of workout thoughts, ideas, and plans.
- **Event Scheduling:** Plan workouts and gym sessions.

## Architecture

The project is structured as a monorepo containing two main components:

*   **Backend (`/api`):** A RESTful API built with Go and the Gin framework, responsible for data management, authentication, and business logic. It uses MongoDB as the primary data store.
*   **Frontend (`/frontend`):** A dynamic web interface built with SvelteKit and TailwindCSS, providing a responsive and interactive user experience.

## Prerequisites

Before running the project, ensure you have the following installed:

*   **Go:** Version 1.25 or higher (for the API).
*   **Bun:** Or Node.js (for the Frontend). Bun is recommended for faster package management.
*   **MongoDB:** A running instance of MongoDB (local or cloud).
*   **Docker:** (Optional) For containerized deployment.

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/AleCro/GYMApplication
cd GYMApplication
```

### 2. Environment Configuration

You will need to set up environment variables for both the API and the Frontend. Refer to the `README.md` in each directory for specific configuration details.

### 3. Running with Docker (Recommended)

The easiest way to get up and running is using Docker Compose.

```bash
docker-compose up --build
```

This command will start both the backend and frontend services, along with a MongoDB instance if configured in the compose file.

### 4. Manual Setup

If you prefer to run the services individually, please refer to the specific documentation for each component:

*   [**Backend Documentation**](./api/README.md)
*   [**Frontend Documentation**](./frontend/README.md)

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

Powered by [yxl-prz/YSvelGoK](https://github.com/yxl-prz/YSvelGoK).
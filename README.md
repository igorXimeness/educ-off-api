# Educ-Off Project

## Overview
Educ-Off is a study app that allows users to download lessons in PDF format for offline viewing. It consists of a **React Native** frontend and a **Go** backend with PostgreSQL, running in a **Docker** environment.

## Requirements

### Frontend (React Native)
To run the **React Native** project, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/igorXimeness/educ-off-api.git
   ```
2. Install Node.js, npm, and Expo.
3. Navigate to the project directory and install dependencies:
   ```sh
   npm install
   ```

### Backend (Go + PostgreSQL)
The **backend** is built with **Go** and runs in a **Docker** environment.

#### Prerequisites
- Install **Go 1.22.0**
- Install **Docker**, **Docker Compose**, and **Docker Desktop**
- Install **DBeaver** (optional, for database management)

#### Running the Backend
1. Clone the repository:
   ```sh
   git clone https://github.com/igorXimeness/educ-off-api.git
   ```
2. Navigate to the project directory:
   ```sh
   cd educ-off-api
   ```
3. Start the services using Docker Compose:
   ```sh
   docker-compose up -d
   ```

### Docker Configuration
#### Dockerfile
```dockerfile
# Base image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy dependency files and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the project files
COPY . .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["go", "run", "cmd/educ-off-api/main.go"]
```

#### docker-compose.yml
```yaml
services:
  db:
    image: postgres:13
    environment:
      POSTGRES_USER: user_teste
      POSTGRES_PASSWORD: password
      POSTGRES_DB: teste
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data # Persistent database storage

  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - db 
    environment:
      DB_HOST: "db"
      DB_USER: "user_teste"
      DB_PASSWORD: "password"
      DB_PORT: "5432"
      DB_NAME: "teste"
    volumes:
      - .:/app # Mounts the application code

volumes:
  pgdata:
```

### Database Access (DBeaver)
To connect to the database using **DBeaver**:
1. Open DBeaver and create a new PostgreSQL connection.
2. Use the following credentials:
   - **Host**: `localhost`
   - **Port**: `5432`
   - **Database**: `teste`
   - **User**: `user_teste`
   - **Password**: `password`


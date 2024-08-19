# EduShare - University Course Resource Management System

EduShare is a platform built with a Vite+React frontend and a Go backend, utilizing a PostgreSQL database. This platform is designed to streamline the process of uploading, verifying, and organizing university course resources.

## Key Features

- **Resource Upload**: Easily upload course materials including previous years' questions (PYQs), books, and other educational resources.
- **Category Sorting**: Automatically categorize uploaded resources for efficient organization and retrieval.
- **Admin Verification**: Ensure the quality and accuracy of resources with an admin verification process before they are made publicly available.
- **User-Friendly Interface**: Intuitive and responsive design for seamless user experience.
- **Secure Storage**: Leveraging PostgreSQL for reliable and secure data storage.

## Motivation

In many colleges, resources are often shared through inefficient methods like WhatsApp or college-provided drives. This project aims to provide a more organized and accessible platform for managing and distributing educational resources.

## Tech Stack

- **Frontend**: Vite, React, shadcn/ui
- **Backend**: Go
- **Database**: PostgreSQL
- **Validation**: zod

## Table of Contents

- [Installation](#installation)
  - [Frontend](#frontend)
  - [Backend](#backend)
- [Usage](#usage)
- [Makefile Details](#makefile-details)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation

### Getting Started

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/edushare.git
    cd edushare
    ```

2. **Set up the PostgreSQL database**:
    - Install PostgreSQL if you haven't already.
    - Create a new database named `edushare`:
      ```bash
      createdb edushare
      ```
    - Configure your PostgreSQL user credentials in the `.env` file.

3. **Configure the backend server**:
    - Create a `.env` file in the `backend` directory with the necessary environment variables (details provided below).

4. **Run the frontend and backend services**:
    - Start the backend server using `make run`.
    - Start the frontend development server using `pnpm run dev`.

### Frontend

1. **Navigate to the frontend directory**:
    ```bash
    cd frontend
    ```

2. **Install the dependencies**:
    ```bash
    pnpm install
    ```

3. **Build the project**:
    ```bash
    pnpm build
    ```

4. **Run the development server**:
    ```bash
    pnpm run dev
    ```

### Backend

1. **Navigate to the backend directory**:
    ```bash
    cd backend
    ```

2. **Install Go on your machine**: [Go installation](https://go.dev/doc/install)

3. **Set up environment variables**: Create a `.env` file in the backend directory (details provided below).

4. **Install the dependencies locally**:
    ```bash
    go mod vendor
    ```

5. **Run database migrations**:
    ```bash
    make migrate-up
    ```

6. **Build the backend**:
    ```bash
    make build
    ```

7. **Run the backend**:
    ```bash
    make run
    ```

## Usage

1. Ensure both the frontend and backend servers are running.

2. Open your browser and go to the frontend development server URL (usually `http://localhost:3000`).

## Makefile Details

The `Makefile` simplifies common tasks during development. Below are the available commands:

- **build**: Compiles the Go backend and places the binary in `bin/build/`.
    ```bash
    make build
    ```
    - Output: `bin/build/edushare` executable.

- **run**: Builds the backend and runs the compiled binary.
    ```bash
    make run
    ```
    - The server will run on the port specified in the `.env` file.

- **clean**: Removes all build artifacts.
    ```bash
    make clean
    ```

- **fmt**: Formats the Go code.
    ```bash
    make fmt
    ```

- **goose-status**: Checks the current status of the database migrations.
    ```bash
    make goose-status
    ```

- **migrate-up**: Applies all pending migrations using Goose.
    ```bash
    make migrate-up
    ```

- **migrate-down**: Rolls back the last applied migration.
    ```bash
    make migrate-down
    ```

## Environment Variables

The backend requires a `.env` file to manage configuration settings. Below is an example:

```bash
PORT=8080
HOST=http://localhost:5173/
DB=postgresql://postgres:postgres@localhost:5432/edushare?sslmode=disable
JWT_SECRET=7*@KMaNg
JWT_EXPIRATION_IN_SECONDS=10080
```
### Environment Variables

- **PORT**: The port on which the backend server will run.
- **HOST**: The frontend URL.
- **DB**: The connection string for the PostgreSQL database.
- **JWT_SECRET**: A secret key used for signing JWT tokens.
- **JWT_EXPIRATION_IN_SECONDS**: The duration for which the JWT token is valid.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch:
    ```bash
    git checkout -b feature-branch
    ```
3. Make your changes and commit them:
    ```bash
    git add .
    git commit -m "Add new feature"
    ```
4. Push to the branch:
    ```bash
    git push origin feature-branch
    ```
5. Create a pull request.

## License

No license.

## Contact

For any questions or suggestions, please contact [abhijith18765@gmail.com](mailto:abhijith18765@gmail.com).

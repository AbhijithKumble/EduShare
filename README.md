# EduShare - University Course Resource Management System

EduShare is a platform built with a Vite+React frontend and a Go backend, utilizing a PostgreSQL database. This platform is designed to streamline the process of uploading, verifying, and organizing university course resources.

Key Features:

- Resource Upload: Easily upload course materials including previous years' questions (PYQs), books, and other educational resources.
- Category Sorting: Automatically categorize uploaded resources for efficient organization and retrieval.
- Admin Verification: Ensure the quality and accuracy of resources with an admin verification process before they are made publicly available.
- User-Friendly Interface: Intuitive and responsive design for seamless user experience.
- Secure Storage: Leveraging PostgreSQL for reliable and secure data storage.

Motivation

In many colleges, resources are often shared through inefficient methods like WhatsApp or college-provided drives. This project aims to provide a more organized and accessible platform for managing and distributing educational resources.
Tech Stack

    Frontend: Vite, React
    Backend: Go
    Database: PostgreSQL

## Table of Contents
- [Installation](#installation)
  - [Frontend](#frontend)
  - [Backend](#backend)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation
    Getting Started
    
      Clone the repository.
      Set up the PostgreSQL database.
      Configure the backend server.
      Run the frontend and backend services.
      
### Frontend

The frontend is built using the Vite framework and React Router 6 for routing. `pnpm` is used as the package manager because it saves space.

1. Navigate to the frontend directory:
    ```bash
    cd frontend
    ```

2. Install the dependencies:
    ```bash
    pnpm install
    ```

3. Build the project:
    ```bash
    pnpm build
    ```

4. Run the development server:
    ```bash
    pnpm run dev
    ```

### Backend

The backend is developed using Go, with Goose for database migration and PostgreSQL as the database.

1. Navigate to the backend directory:
    ```bash
    cd backend
    ```

2. Install Go on your machine: [Go installation](https://go.dev/doc/install)

3. Check the Makefile for easier build and install commands.

4. To migrate databases specified in the Makefile, use:
    ```bash
    make migrate-up
    make migrate-down
    ```

5. Install the dependencies locally:
    ```bash
    go mod vendor
    ```

6. Install all dependencies in the project directory:
    ```bash
    go get .
    ```

7. Clean all unused dependencies:
    ```bash
    go mod tidy
    ```

8. Install all dependencies to the root Go folder:
    ```bash
    go install .
    ```

9. Set the environment variables as required.

10. Build the project into the output directory:
    ```bash
    go build -o build/edushare/
    ```

## Usage

1. Ensure both the frontend and backend servers are running.

2. Open your browser and go to the frontend development server URL (usually `http://localhost:3000`).

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
   - Use clear and concise commit messages.
   - Follow the commit message convention:
     - **feat:** A new feature
     - **fix:** A bug fix
     - **docs:** Documentation only changes
     - **style:** Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
     - **refactor:** A code change that neither fixes a bug nor adds a feature
     - **perf:** A code change that improves performance
     - **test:** Adding missing or correcting existing tests
     - **chore:** Changes to the build process or auxiliary tools and libraries such as documentation generation
   - Example commit message:
     ```bash
     git commit -m "feat: add user authentication"
     ```
4. Push to the branch:
    ```bash
    git push origin feature-branch
    ```
5. Create a pull request.
6. Make sure to add test cases to the project.

## License

No license

## Contact

For any questions or suggestions, please contact [abhijith18765@gmail.com](mailto:abhijith18765@gmail.com).

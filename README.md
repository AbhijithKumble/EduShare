# EduShare

EduShare is a platform that stores resources to study various courses offered by a University.

## Table of Contents
- [Installation](#installation)
  - [Frontend](#frontend)
  - [Backend](#backend)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation

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
    git commit -m "Add new feature"
    ```
4. Push to the branch:
    ```bash
    git push origin feature-branch
    ```
5. Create a pull request.

## License

No license

## Contact

For any questions or suggestions, please contact [abhijith18765@gmail.com](mailto:abhijith18765@gmail.com).

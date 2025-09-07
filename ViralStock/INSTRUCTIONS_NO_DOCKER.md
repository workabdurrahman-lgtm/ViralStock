# Instructions for Running ViralStock without Docker

This guide provides instructions on how to run the ViralStock frontend and backend services on your local machine without using Docker.

## Prerequisites

Before you begin, you will need to have the following installed on your system:

- **Node.js**: Version 18 or a compatible version. You can download it from [https://nodejs.org/](https://nodejs.org/).
- **Go**: Version 1.20 or a compatible version. You can download it from [https://go.dev/](https://go.dev/).

You can verify the installations by running `node -v` and `go version` in your terminal.

## Running the Frontend

The frontend is a Svelte application. To run it, follow these steps:

1.  **Navigate to the frontend directory:**
    ```bash
    cd ViralStock/frontend
    ```

2.  **Install the necessary Node.js packages:**
    ```bash
    npm install
    ```

3.  **Start the development server:**
    ```bash
    npm run dev
    ```

    The frontend will now be running and accessible at `http://localhost:5173`.

## Running the Backend

The backend is a Go application. To run it, follow these steps:

1.  **Navigate to the backend directory:**
    ```bash
    cd ViralStock/backend
    ```

2.  **Download the Go modules (dependencies):**
    This command will download any new dependencies that have been added to the project.
    ```bash
    go mod download
    ```

3.  **Build the application:**
    ```bash
    go build -o main .
    ```

4.  **Run the compiled server:**
    ```bash
    ./main
    ```

    The backend will now be running and accessible at `http://localhost:8080`.

## Running Both Services

To run the entire application, you will need to open two separate terminal windows or tabs: one for the frontend and one for the backend. Follow the respective instructions in each terminal. Once both are running, the application will be fully functional.

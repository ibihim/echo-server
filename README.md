# Golang Echo Server

This is a simple Golang echo server that listens for HTTP requests and echoes back the exact request received, including the HTTP protocol, headers, and body.

## Features

- Echoes back the exact HTTP request received
- Logs incoming requests and dumped request details
- Containerized using Docker for easy deployment

## Prerequisites

- Go (version 1.22 or later)
- Docker (for containerized deployment)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/ibihim/echo-server
```

2. Change into the project directory:

```bash
cd echo-server
```

3. Run the server:

```bash
go run main.go
```

The server will start running on `http://localhost:8080`.

4. Send HTTP requests to the server using tools like cURL or a web browser.

Example cURL command:

```bash
curl -X POST -H "Content-Type: text/plain" \
    -H "Authorization: Bearer fake-token" \
    -d "hello world" http://localhost:8080
```

The server will echo back the exact request received.

## Containerized Deployment

To deploy the echo server using Docker, follow these steps:

1. Build the Docker image:

```bash
docker build -t echo-server .
```

2. Run the Docker container:

```bash
docker run -p 8080:8080 echo-server
```

The server will start running inside the container and will be accessible on `http://localhost:8080`.

## Testing

To run the unit tests for the echo server, use the following command:

```bash
go test
```

The tests will make a request to the server and verify that the response body contains the expected content based on the request.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

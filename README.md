
# Stored XSS Testing Tool

This standalone tool is designed to test for **Stored Cross-Site Scripting (XSS)** vulnerabilities on web applications. It injects XSS payloads into a target endpoint and checks for the persistence of the payloads in subsequent requests.

## Features

- **Automated Testing**: Tests various XSS payloads for persistence.
- **Dockerized Environment**: Easily deployable in a containerized setup (optional).
- **Logs HTTP Responses**: Monitors the status of each request to track vulnerabilities.

## Usage

### Step 1: Clone or Copy the Repository
Download the tool to your local environment:
```bash
git clone https://github.com/YamatoGr93/stored-xxs.git
cd stored-xxs
```

### Step 2: Build the Docker Image (Optional)
If you want to run the tool in a Docker container:
```bash
docker build -t stored-xss .
```

### Step 3: Run the Tool
Run the tool against a target endpoint. Replace `<endpoint>` with the vulnerable path (e.g., `/comment`):
```bash
docker run --rm --network host stored-xss "<endpoint>"
```

Example:
```bash
docker run --rm --network host stored-xss "/comment"
```

### Expected Output
The tool will generate test URLs with payloads and log the HTTP responses. Example output:
```plaintext
Testing Stored XSS on endpoint: http://localhost:3000/comment

Tested: http://localhost:3000/comment<script>alert('Stored XSS')</script> | Status: 200
Tested: http://localhost:3000/comment"><script>alert('Stored XSS')</script> | Status: 200
```
## Troubleshooting

### Common Issues

1. **Connection Refused**: Ensure the target application is running and accessible at the specified URL.
2. **Host Resolution Errors**: Ensure Docker is running in `--network host` mode for local connections.

## Disclaimer

This tool is for **educational and authorized penetration testing purposes only**. Use responsibly on systems you have permission to test.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

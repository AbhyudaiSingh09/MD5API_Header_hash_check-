

# Binary File MD5 Hash Checker

This Go application, utilizing the Gin framework, provides a web interface that allows users to upload binary files. It specifically extracts the header information, calculates the MD5 hash of the file's content excluding the header, and compares it to an expected hash value specified within the file's header.

## Installation

Ensure you have Go installed on your system to run this application. Follow these installation steps:

1. Clone the repository or download the source code.
2. Navigate to the application's directory.
3. Execute `go mod tidy` to install the required dependencies.

## Running the Application

Start the application with the command:

```sh
go run main.go
```

This command initiates the web server on port `8080`. Access the file upload interface by navigating to `http://localhost:8080/` in your web browser.

## Usage

- Open `http://localhost:8080/` in a web browser to access the upload interface.
- Use the form to select a binary file and upload it.
- The application will extract the header from the binary file, calculate the MD5 hash of the content following the header, and compare this calculated hash with the expected hash value defined within the header.
- The response will indicate whether the calculated hash matches the expected hash found in the file's header.

## Test File

The application includes a test file named `main_test.go` which contains unit tests. Execute these tests using the following command:

```sh
go test -v
```

## Video Tutorial

For an in-depth walkthrough and demonstration, please refer to the video tutorial linked below:

[![Binary File MD5 Hash Checker](http://img.youtube.com/vi/VIDEO_ID/0.jpg)](http://www.youtube.com/watch?v=VIDEO_ID "Binary File MD5 Hash Checker")

[View Demo Video](https://github.com/Axs7941/MD5API_Header_hash_check-/blob/main/MD5api_Demo.mov)


package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFileUploadWithActualFile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	setupRouter(router) // Adjust this to your actual router setup function

	// The path to the file you want to upload
	filePath := "path/to/your/testfile.txt" // Update this to the path of your test file

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a buffer to hold the multipart form data
	var requestBody bytes.Buffer
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Create a form field for the file
	fileWriter, err := multiPartWriter.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Errorf("Error creating form file: %v", err)
	}

	// Copy the file content to the form field writer
	if _, err := io.Copy(fileWriter, file); err != nil {
		t.Errorf("Error writing to form file: %v", err)
	}

	// Important to close the writer to finalize the form data
	multiPartWriter.Close()

	// Create a new request with the form data
	req, _ := http.NewRequest(http.MethodPost, "/upload", &requestBody)
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code, "Expected response code to be 200 OK")

	// Here you can add more assertions based on the expected outcome
	// For example, checking the response body for specific result messages
	assert.Contains(t, w.Body.String(), "Expected hash matches the calculated hash.", "Expected success message in response body")
	// Or adjust the assertion based on your application's actual response
}

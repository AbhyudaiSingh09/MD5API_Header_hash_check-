package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 10 << 20 // 10 MiB
	router.GET("/", func(c *gin.Context) {
		c.File("upload.html") // Ensure this file exists in the same directory as your main.go
	})

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
			return
		}

		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer openedFile.Close()

		content, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
			return
		}

		startMarker := []byte("[[")
		endMarker := []byte("]]")

		startIndex := indexOf(content, startMarker) + len(startMarker)
		endIndex := indexOf(content, endMarker)

		if startIndex == -1 || endIndex == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format: markers not found"})
			return
		}

		// Extract the header to get the third element
		header := content[startIndex-2 : endIndex+2]
		// Convert header to a string for manipulation
		headerStr := string(header)
		// Trim the surrounding brackets
		headerStr = headerStr[2 : len(headerStr)-2]
		// Split the string by "," to access individual elements
		elements := strings.Split(headerStr, ",")

		// Check if we have a third element
		var expectedHash string
		if len(elements) >= 3 {
			expectedHash = elements[2]
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Header does not contain a third element"})
			return
		}

		// Proceed to hash the content excluding the header
		dataToHash := content[endIndex+2 : len(content)-3] // Assuming the content ends right after the header
		calculatedHash := md5.Sum(dataToHash)
		calculatedHashString := hex.EncodeToString(calculatedHash[:])
		calculatedHashString = strings.ToUpper(calculatedHashString) // Convert calculated hash to uppercase

		// Check if the expected hash matches the calculated hash
		if expectedHash == calculatedHashString {
			c.JSON(http.StatusOK, gin.H{
				"match": true,
				"message": "Expected hash matches the calculated hash.",
				"ExpectedHash": expectedHash,
				"CalculatedHash": calculatedHashString,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"match": false,
				"message": "Expected hash does not match the calculated hash.",
				"ExpectedHash": expectedHash,
				"CalculatedHash": calculatedHashString,
			})
		}
	})

	router.Run(":8080")
}

// indexOf finds the index of a sequence within a byte slice and returns -1 if not found.
func indexOf(slice []byte, toFind []byte) int {
	for i := 0; i < len(slice)-len(toFind)+1; i++ {
		if bytes.Equal(slice[i:i+len(toFind)], toFind) {
			return i
		}
	}
	return -1
}

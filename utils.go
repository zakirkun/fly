package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Check if a folder exists, if not, create it
func ensureFolderExists(folderPath string) error {
	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create the folder with appropriate permissions (e.g., 0755)
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create folder: %v", err)
		}
		log.Println("Folder created:", folderPath)
	}
	return nil
}

// Check if a file exists, if not, create it
func ensureFileExists(filePath string) error {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()

		log.Println("File created:", filePath)
	}
	return nil
}

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

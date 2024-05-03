package main

import (
	"SPADE/usecases/hypnogram"
	"SPADE/utils"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestUser(t *testing.T) {
	// read the files from dataset one file per user
	//
	// read the user's data from file
	dir := "../dataset/"
	//	fileName := "b000101.txt"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for i, fileInfo := range files {
		// Check if the file is a .txt file
		if filepath.Ext(fileInfo.Name()) == ".txt" {
			// Open the file
			filePath := filepath.Join(dir, fileInfo.Name())
			data := utils.AddPadding(hypnogram.PaddingItem, hypnogram.MaxVecSize, utils.ReadFile(filePath))
			testUser(t, i, data)
		}
	}
}

func testUser(t *testing.T, id int, data []int) {
	fmt.Printf("=== User Test: id = %d, data lenght=%d\n", id, len(data))

	t.Run("RunUser", func(t *testing.T) {
		// check if you should add delay here for server connections
		RunUser(id, data)
	})
}

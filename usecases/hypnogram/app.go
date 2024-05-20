package hypnogram

import (
	"SPADE/usecases/models"
	"SPADE/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// OpenDataset read the hypnogram files from dataset directory
// assumption: one file per each user
func OpenDataset() [][]int {
	fmt.Println("Opening dataset..")
	dir := "./dataset/"
	//	fileName := "b000101.txt"
	files, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println("Error reading dataset directory:", err)
		return nil
	}

	fmt.Printf("Reading %d files\n", len(files))

	dataset := make([][]int, 0)
	for _, fileInfo := range files {
		// Check if the file is a .txt file
		if filepath.Ext(fileInfo.Name()) == ".txt" {
			// Open the file
			filePath := filepath.Join(dir, fileInfo.Name())
			data := utils.AddPadding(PaddingItem, MaxVecSize, utils.ReadHypnogramFile(filePath))
			dataset = append(dataset, data)
		}
	}
	return dataset
}

// InitServer start the SPADE server using hypnogram use-case configuration
func InitServer() {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	models.StartServer(config)
}

// InitUsers start the SPADE users using hypnogram use-case configuration
func InitUsers() {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	dataset := OpenDataset()
	for i, hypnogram := range dataset {
		fmt.Printf("=== User Test: id = %d, data length=%d\n", i, len(hypnogram))
		models.StartUser(config, i, hypnogram)
	}
}

// InitAnalyst start the SPADE analyst using hypnogram use-case configuration
func InitAnalyst() {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	userID := int64(0)
	queryValue := int64(3)
	// start analyst entity, send a req to server for getting the cipher belongs to
	// the user with userID and decrypt it for the specific query value queryValue
	queryValue, results := models.StartAnalyst(config, userID, queryValue)
	// !!! WARNING !!!!
	// THIS IS TO VERIFY THE results and proof the correctness of the protocol
	// we are not going to do this in a real world application, Hopefully :)
	// read the original user's data from file and compare it with the results
	datasetDir := "./dataset/"
	fileName := "b000101.txt"
	data := utils.AddPadding(config.PaddingItem, config.MaxVecSize, utils.ReadHypnogramFile(datasetDir+fileName))
	log.Println("Decrypted Data:", results)
	utils.VerifyResults(data, results, int(queryValue))
	log.Println(">>> Analyst's operations are done!")
}

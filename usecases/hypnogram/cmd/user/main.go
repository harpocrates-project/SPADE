// package main
// 
// import (
// 	"SPADE/usecases/models"
// 	"SPADE/utils"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )
// 
// // OpenDataset read the hypnogram files from dataset directory
// // assumption: one file per each user
// func OpenDataset() [][]int {
// 	fmt.Println("Opening dataset..")
// 	dir := "./dataset/"
// 	//	fileName := "b000101.txt"
// 	files, err := os.ReadDir(dir)
// 
// 	if err != nil {
// 		fmt.Println("Error reading dataset directory:", err)
// 		return nil
// 	}
// 
// 	fmt.Printf("Reading %d files\n", len(files))
// 
// 	dataset := make([][]int, 0)
// 	for _, fileInfo := range files {
// 		// Check if the file is a .txt file
// 		if filepath.Ext(fileInfo.Name()) == ".txt" {
// 			// Open the file
// 			filePath := filepath.Join(dir, fileInfo.Name())
// 			hypnogram := utils.ReadHypnogramFile(filePath)
// 			data := utils.AddPadding(PaddingItem, MaxVecSize, hypnogram)
// 			dataset = append(dataset, data)
// 			// for manual testing, I am just reading a single file from dataset
// 			// to use the full test, please comment the following break
// 			break
// 		}
// 	}
// 	return dataset
// }
// 
// // InitUsers start the SPADE users using hypnogram use-case configuration
// func InitUsers() {
// 	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
// 	dataset := OpenDataset()
// 	for i, hypnogram := range dataset {
// 		fmt.Printf("=== User Test: id = %d, data length=%d\n", i, len(hypnogram))
// 		models.StartUser(config, i, hypnogram)
// 	}
// }
// 
// 
// func main() {
// 	InitUsers()
// }

package main

import (
	"SPADE/usecases/models"
	"SPADE/utils"
	"fmt"
	"strconv"
	"os"
)

// Upload hynogram at filePath with id
func InitUser(serverAddr string, id int, filePath string) {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	hypnogram := utils.ReadHypnogramFile(filePath)
	data := utils.AddPadding(PaddingItem, MaxVecSize, hypnogram)
	models.StartUser(serverAddr, config, id, data)
}

func main() {
	serverAddr := os.Args[1]
    	id, err := strconv.Atoi(os.Args[2])
    	if err != nil {
		fmt.Println("Error converting id to integer", err)
		panic(err)
    	}
	filePath := os.Args[3]
	InitUser(serverAddr, id, filePath)
}

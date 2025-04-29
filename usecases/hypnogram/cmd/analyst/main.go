package main

import (
	"SPADE/usecases/models"
	"log"
	"os"
        "fmt"
        "strconv"
)

// InitAnalyst start the SPADE analyst using hypnogram use-case configuration
func InitAnalyst(serverAddr string, userID int64, queryValue int64) {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	// start analyst entity, send a req to server for getting the cipher belongs to
	// the user with userID and decrypt it for the specific query value queryValue
	queryValue, results := models.StartAnalyst(serverAddr, config, userID, queryValue)

	// !!! WARNING !!!!
	// THIS IS TO VERIFY THE results and proof the correctness of the protocol
	// we are not going to do this in a real world application, Hopefully :)
	// read the original user's data from file and compare it with the results
	// datasetDir := "./dataset/"
	// fileName := "b000101.txt"
	//data := utils.AddPadding(config.PaddingItem, config.MaxVecSize, utils.ReadHypnogramFile(datasetDir+fileName))
	// utils.VerifyResults(data, results, int(queryValue))
	// log.Println(">>> Analyst's operations are done!")

	log.Println("Decrypted Data:", results)
}

func main() {
	serverAddr := os.Args[1]
	userID, err := strconv.ParseInt(os.Args[2], 10, 64)
    	if err != nil {
		fmt.Println("Error converting user id to 64bit integer", err)
		panic(err)
    	}
	// Why 3?
	InitAnalyst(serverAddr, userID, 3)
}

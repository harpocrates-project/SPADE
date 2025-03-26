package main

import (
	"SPADE/usecases/models"
	"log"
)

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
	// datasetDir := "./dataset/"
	// fileName := "b000101.txt"
	//data := utils.AddPadding(config.PaddingItem, config.MaxVecSize, utils.ReadHypnogramFile(datasetDir+fileName))
	// utils.VerifyResults(data, results, int(queryValue))
	// log.Println(">>> Analyst's operations are done!")

	log.Println("Decrypted Data:", results)
}

func main() {
	InitAnalyst()
}

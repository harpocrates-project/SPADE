package main

import (
	"SPADE/usecases/models"
	"os"
)

// InitServer start the SPADE server using hypnogram use-case configuration
func InitServer(serverAddr string) {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	models.StartServer(serverAddr, config)
}

func main() {
	serverAddr := os.Args[1]
	InitServer(serverAddr)
}

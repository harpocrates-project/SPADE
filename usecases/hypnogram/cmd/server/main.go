package main

import (
	"SPADE/usecases/models"
)

// InitServer start the SPADE server using hypnogram use-case configuration
func InitServer() {
	config := models.NewConfig(DbName, TbName, NumUsers, MaxVecSize, PaddingItem, TimeOut, MaxMsgSize)
	models.StartServer(config)
}

func main() {
	InitServer()
}

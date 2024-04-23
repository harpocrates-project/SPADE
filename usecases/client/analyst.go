package main

import (
	"SPADE"
	pb "SPADE/spadeProto"
	"SPADE/usecases"
	"SPADE/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/big"
	"time"
)

type Analyst struct {
	id    int
	q     *big.Int
	g     *big.Int
	spade *SPADE.SPADE
	mpk   []*big.Int
}

func NewAnalyst(q, g *big.Int, mpk []*big.Int) *Analyst {
	return &Analyst{
		id:    1,
		q:     q,
		g:     g,
		spade: nil,
		mpk:   mpk,
	}
}

func main() {
	pbHandler := usecases.NewPBHandler()

	log.Println(">>> Analyst starts connecting to the server..")
	addr := fmt.Sprintf("localhost:%d", utils.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// proto buffer init
	a := pb.NewCuratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// SPADE calls for Analysts here :)
	q, g, mpk, err := pbHandler.ReadPublicParams(a.GetPublicParams(ctx, &pb.PublicParamsReq{}))
	if err != nil {
		log.Fatalf("could not fetch the public parameters: %v", err)
	}

	// create a new Analyst
	// create an instance of SPADE with same public params of server
	analyst := NewAnalyst(q, g, mpk)
	spd := SPADE.NewSpade(q, g, usecases.MaxVecSize)
	analyst.spade = spd

	utils.PrintBigIntHex("q", q)
	utils.PrintBigIntHex("g", g)

	// send a query for value(1) and user-id(1) to the server
	req := &pb.AnalystReq{
		Id:    1,
		Value: 3,
	}

	// get the unmarshal values
	dkv, cts, err := pbHandler.ReadDecryptionKey(a.Query(ctx, req))
	if err != nil {
		log.Fatalf("could not send the request: %v", err)
	}

	// partially decrypt the ciphertext vector using dkv to get the result for query
	results := spd.Decrypt(dkv, int(req.Value), cts)

	// !!! WARNING !!!!
	// THIS IS TO VERIFY THE results and proof the correctness of the protocol
	// we are not going to do this in a real world application, Hopefully :)
	// read the original user's data from file
	datasetDir := "../dataset/"
	fileName := "b000101.txt"
	data := utils.AddPadding(usecases.PaddingItem, usecases.MaxVecSize, utils.ReadFile(datasetDir+fileName))
	log.Println(results)
	utils.VerifyResults(data, results, int(req.Value))
	log.Println(">>> Analyst's operations are done!")
}

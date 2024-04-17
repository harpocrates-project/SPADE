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

type User struct {
	id    int
	q     *big.Int
	g     *big.Int
	alpha *big.Int
	mpk   []*big.Int
}

func NewUser(q, g *big.Int, mpk []*big.Int) *User {
	return &User{
		id:    15,
		q:     q,
		g:     g,
		alpha: nil,
		mpk:   mpk,
	}
}

func main() {
	pbHandler := usecases.NewPBHandler()

	log.Println(">>> Client starts connecting to the server..")
	addr := fmt.Sprintf("localhost:%d", utils.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// proto buffer init
	c := pb.NewCuratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// SPADE calls for Users here :)
	q, g, mpk, err := pbHandler.ReadPublicParams(c.GetPublicParams(ctx, &pb.PublicParamsReq{}))
	if err != nil {
		log.Fatalf("could not fetch the public parameters: %v", err)
	}

	// create a new user
	// create an instance of SPADE with same public params of server
	// generate a random secret for the user
	user := NewUser(q, g, mpk)
	spade := SPADE.NewSpade(q, g, usecases.MaxVecSize)
	user.alpha = SPADE.RandomElementInZMod(q)
	regKey := spade.Register(user.alpha)

	utils.PrintBigIntHex("q", q)
	utils.PrintBigIntHex("g", g)
	utils.PrintBigIntHex("regKey", regKey)

	// read the user's data from file
	datasetDir := "../dataset/"
	fileName := "b000101.txt"
	data := utils.AddPadding(usecases.PaddingItem, usecases.MaxVecSize, utils.ReadFile(datasetDir+fileName))

	// encrypt user's data using mpk
	ct := spade.Encrypt(user.mpk, user.alpha, data)
	ctBytes := make([][]byte, 0, len(ct)) // Pre-allocate for efficiency
	for _, row := range ct {
		for _, item := range row {
			ctBytes = append(ctBytes, item.Bytes())
		}
	}

	// send the encrypted data to the server
	encData := &pb.UserReq{
		Id:         int64(user.id),
		RegKey:     regKey.Bytes(),
		Ciphertext: ctBytes,
	}
	ack, err := c.UserRequest(ctx, encData)
	if err != nil {
		log.Fatalf("could not send user data: %v", err)
	}

	if ack.Flag {
		log.Printf("User ACK flag is true!")
	} else {
		log.Printf("There is a problem with the user! Kill Bill")
	}

	log.Println(">>> User's operations are done!")
}

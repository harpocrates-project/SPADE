package main

import (
	"SPADE"
	pb "SPADE/spadeProto"
	"SPADE/usecases/hypnogram"
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

func NewUser(uid int, q, g *big.Int, mpk []*big.Int) *User {
	return &User{
		id:    uid,
		q:     q,
		g:     g,
		alpha: nil,
		mpk:   mpk,
	}
}

func RunUser(id int, data []int) (u *User) {
	start := time.Now()
	pbHandler := hypnogram.NewPBHandler()

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
	user := NewUser(id, q, g, mpk)
	spade := SPADE.NewSpade(q, g, hypnogram.MaxVecSize)
	user.alpha = SPADE.RandomElementInZMod(q)
	regKey := spade.Register(user.alpha)

	utils.PrintBigIntHex("q", q)
	utils.PrintBigIntHex("g", g)
	utils.PrintBigIntHex("regKey", regKey)

	// encrypt user's data using "mpk"
	ct := spade.Encrypt(user.mpk, user.alpha, data)
	// log.Println(">>> ct[0]: ", ct[0])
	// utils.SaveInFile("./ciphertext.txt", ct)

	// Note: here we encode the ct = [n][2]*big.Int into ctBytes = [n*2][t]byte,
	// where t=len(ct_element.Bytes()), i.e. here will be ctBytes = [n*2][16]byte.
	ctBytes := make([][]byte, 0, len(ct))
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

	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("User finished in %s", elapsed)

	log.Println(">>> User's operations are done!")

	return user
}

//func main() {
//	log.Println(">>> Client is starting...")
//	// read the user's data from file
//	datasetDir := "../dataset/"
//	fileName := "b000101.txt"
//	data := utils.AddPadding(hypnogram.PaddingItem, hypnogram.MaxVecSize, utils.ReadFile(datasetDir+fileName))
//	RunUser(591, data)
//}

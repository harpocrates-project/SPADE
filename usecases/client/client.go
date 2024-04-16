package main

import (
	"SPADE"
	pb "SPADE/spadeProto"
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
	spade *SPADE.SPADE
	mpk   []*big.Int
}

func NewUser(q, g *big.Int, mpk []*big.Int) *User {
	return &User{
		id:    1,
		q:     q,
		g:     g,
		alpha: nil,
		spade: nil,
		mpk:   mpk,
	}
}

func main() {
	addr := fmt.Sprintf("localhost:%d", utils.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCuratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// SPADE calls here :)
	q, g, mpk, err := readPublicParams(c.GetPublicParams(ctx, &pb.PublicParamsReq{}))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	user := NewUser(q, g, mpk)
	spd := SPADE.NewSpade(q, g)
	user.spade = spd
	user.alpha = SPADE.RandomElementInZMod(q)
	regKey := spd.Register(user.alpha)

	utils.PrintBigIntHex("q", q)
	utils.PrintBigIntHex("g", g)
	utils.PrintBigIntHex("regKey", regKey)
}

// readPublicParams convert the byte stream data from server into required data type for client
func readPublicParams(res *pb.PublicParamsRes, err error) (*big.Int, *big.Int, []*big.Int, error) {
	if err != nil {
		return nil, nil, nil, err
	}

	// Unmarshal q, g
	q := new(big.Int)
	q.SetBytes(res.Q)

	g := new(big.Int)
	g.SetBytes(res.G)

	// Unmarshal mpk (slice of big.Int)
	mpk := make([]*big.Int, 0, len(res.Mpk))
	for _, mpkBytes := range res.Mpk {
		temp := new(big.Int)
		temp.SetBytes(mpkBytes)
		mpk = append(mpk, temp)
	}

	return q, g, mpk, nil
}

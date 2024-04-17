package usecases

import (
	pb "SPADE/spadeProto"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"
	"math/big"
)

type PBHandler interface {
	ReadPublicParams(res *pb.PublicParamsResp, err error) (*big.Int, *big.Int, []*big.Int, error)
}

type pbHandler struct{}

func NewPBHandler() PBHandler {
	return &pbHandler{}
}

// ReadPublicParams convert the byte stream data from server into required data type for client
func (pbh pbHandler) ReadPublicParams(response *pb.PublicParamsResp, err error) (*big.Int, *big.Int, []*big.Int, error) {
	if err != nil {
		return nil, nil, nil, err
	}

	// Unmarshal q, g
	q := new(big.Int)
	q.SetBytes(response.Q)

	g := new(big.Int)
	g.SetBytes(response.G)

	// Unmarshal mpk (slice of big.Int)
	mpk := make([]*big.Int, 0, len(response.Mpk))
	for _, mpkBytes := range response.Mpk {
		temp := new(big.Int)
		temp.SetBytes(mpkBytes)
		mpk = append(mpk, temp)
	}

	return q, g, mpk, nil
}

// DBHandler is an interface to work with sqlite3 database API
type DBHandler interface {
	CreateUsersCipherTable() error
	InsertUsersCipher(data *pb.UserReq) error
}

type dbHandler struct {
	DbName string
	TbName string
}

func (d dbHandler) InsertUsersCipher(data *pb.UserReq) error {
	// Marshal user data to JSON bytes
	row, err := proto.Marshal(data)
	if err != nil {
		return err
	}

	// open the database
	db, err := sql.Open("sqlite3", d.DbName)
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert the data into the table
	insertQuery := "INSERT INTO " + d.TbName + " (id, regKey, ciphertext) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, data.Id, row, row)
	if err != nil {
		return err
	}

	// OK
	return nil
}

func (d dbHandler) CreateUsersCipherTable() error {
	// open the database
	db, err := sql.Open("sqlite3", d.DbName)
	if err != nil {
		return err
	}
	defer db.Close()

	// create the table if it doesn't exist (schema migration)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ` + d.TbName + ` (
		id INTEGER PRIMARY KEY,
		regKey BLOB,
		ciphertext BLOB
	)`)
	if err != nil {
		return err
	}

	// OK
	return nil
}

func NewDBHandler() DBHandler {
	return &dbHandler{
		DbName: DbName,
		TbName: TbName,
	}
}

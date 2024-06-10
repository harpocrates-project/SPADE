# SPADE: **S**elective and **PA**rtial **DE**cryption using Functional Encryption

SPADE is an FE-based scheme that enables running selective and partial decryption queries over
vectors of ciphertexts.

To see two real-world applications of SPADE please check [this](https://github.com/harpocrates-project/SPADE/blob/main/usecases/README.md).

## Notes

1. Curators are required to specify both the number of users and the maximum number of entries
   allowed for each plaintext vector. Users must encrypt their data as a vector of integers,
   ensuring that none of the values are **zero**.
2. Be careful when defining the users' data vector using `make()`;
   this method assigns the **zero** value to the elements during initialization.

## Changing the protobuf structure

If you want to modify the protobuf structure, please first change the following file:

    ./spadeproto/spade.proto

and then run the proto compiler command as follows to generate the new protobuf files:

    protoc --go_out=. --go-grpc_out=. spade.proto 

## Testing Instruction

    go test     

## Benchmarking Instruction

    go test -benchtime=10x -bench=BenchmarkSpade -benchmem -run=^$


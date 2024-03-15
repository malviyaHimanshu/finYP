package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	audit "finYP/cmd/contract/abigen"
	functions "finYP/utils/func"

	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CommContract(paramsByte []byte) {
	gethPath, nodeKey := functions.GethPathAndKey()

	connection, err := ethclient.Dial(gethPath)
	functions.CheckError(err)
	fmt.Println("connection Returned: ", connection)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(nodeKey), "asdfgh", big.NewInt(1907))
	fmt.Println("Auth : ", auth)
	functions.CheckError(err)

	addr, tx, instContract, err := audit.DeployAudit(auth, connection)
	functions.CheckError(err)
	fmt.Println("addr Returned: ", addr)
	fmt.Println("instContract Returned: ", instContract)

	// Write contract address to file for other app reference
	functions.FileCrWr("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/contractInfo/contract-address", string([]byte(addr.Hex())))

	ctx := context.Background()
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)

	pairing, err := pbc.NewPairingFromString(string(paramsByte))
	functions.CheckError(err)
	g1 := pairing.NewG1().Rand()
	g2 := pairing.NewG2().Rand()
	//u := pairing.NewG1().Rand()

	tx1, err := instContract.SetParamsGU(auth, paramsByte, g1.Bytes(), g2.Bytes())
	functions.CheckError(err)
	//fmt.Println("\nTxn1 hash is: ", tx1.Hash())
	_, err = bind.WaitMined(ctx, connection, tx1)
	functions.CheckError(err)
	fmt.Println("Contract Deployment Successful!")
}

func main() {
	// create ./tmp/ folder for storing secrets information of parties and public parameters
	functions.Createtmp()

	// In a real application, generate this once and publish it
	params := pbc.GenerateA(160, 512)

	CommContract([]byte(params.String()))

}

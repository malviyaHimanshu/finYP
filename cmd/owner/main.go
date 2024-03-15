package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/csv"
	audit "finYP/cmd/contract/abigen"
	functions "finYP/utils/func"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var pairing *pbc.Pairing
var g1 *pbc.Element
var g2 *pbc.Element
var pk *pbc.Element
var skt *pbc.Element
var skh *pbc.Element
var x []*pbc.Element

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("function %s took %s", name, elapsed)
}

func err_message() {
	fmt.Printf("\nWelcome to Owner(client) program...\n")
	fmt.Printf("format: $ go build owner_n.go\n")
	fmt.Printf("\t$ ./owner_n PATH_TO_FILENAME\n\n")
	os.Exit(8)
}

func ContractVars() (*audit.Audit, *bind.TransactOpts, *ethclient.Client) {
	gethPath, nodeKey := functions.GethPathAndKey()

	addrByte, err := ioutil.ReadFile("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/contractInfo/contract-address")
	functions.CheckError(err)

	connection, err := ethclient.Dial(gethPath)
	functions.CheckError(err)
	fmt.Println("connection Returned: ", connection)

	instContract, err := audit.NewAudit(common.HexToAddress(string(addrByte)), connection)
	functions.CheckError(err)
	fmt.Println("instContract Returned: ", instContract)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(nodeKey), "asdfgh", big.NewInt(1907))
	functions.CheckError(err)

	return instContract, auth, connection
}

func main() {
	instContract, auth, connection := ContractVars()
	ctx := context.Background()

	paramsByte, GetGVal, GetUVal, err := instContract.GetParamsGU(nil)
	functions.CheckError(err)

	//fmt.Println("check!!!!")
	pairing, err = pbc.NewPairingFromString(string(paramsByte))
	functions.CheckError(err)

	g1 = pairing.NewG1().SetBytes(GetGVal)
	g2 = pairing.NewG2().SetBytes(GetUVal)
	skt = pairing.NewZr().Rand()
	skh = pairing.NewZr().Rand()
	pk = pairing.NewG2().PowZn(g2, skt)

	fmt.Println("g1 : ", g1)
	fmt.Println("g2 : ", g2)
	fmt.Println("skh : ", skh.Bytes())

	functions.FileCrWr("/home/ajay/Desktop/src/contract/tmp/private/owner/skt", string(skt.Bytes()))
	functions.FileCrWr("/home/ajay/Desktop/src/contract/tmp/private/owner/skh", string(skh.Bytes()))

	tx, err := instContract.SetOwnerKeys(auth, pk.Bytes(), skh.Bytes())
	functions.CheckError(err)
	fmt.Println("\nTxn hash is: ", tx.Hash())
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)

	GetPKOVal, err := instContract.GetPKT(nil)
	functions.CheckError(err)
	fmt.Println("\nGetPKO() Returned: ", GetPKOVal)

	msg := "\nconnecting to the server http://127.0.0.1:8003/...\n"
	msgtype := ""
	msgloc := "http://127.0.0.1:8003/"
	ret := functions.Connect(msgloc, msg, msgtype)
	fmt.Println(ret)

	if len(os.Args) < 2 {
		err_message()
	}

	file, err := os.Open(os.Args[1])
	functions.CheckError(err)
	defer file.Close()

	data := make([]byte, pairing.ZrLength()-1)
	noOfBlocks := 0 //Also number of data points
	NumOfSectors := functions.TotalSectors()
	var EndOfFile bool = false

	x := func() {
		var filearr [][]*pbc.Element

		for !EndOfFile {
			noOfBlocks += 1
			var block []*pbc.Element

			for i := 0; i < NumOfSectors; i++ {
				_, err := file.Read(data)
				if err != nil {
					if err != io.EOF {
						log.Fatal(err)
					}
					EndOfFile = true
					break
				}

				m := pairing.NewZr().SetBytes(data)
				block = append(block, m)
			}

			filearr = append(filearr, block)
		}

		s_n := strconv.Itoa(noOfBlocks)
		for i := 0; i < NumOfSectors; i++ {
			x = append(x, pairing.NewZr().Rand())
		}

		start := time.Now()

		for i := 0; i < noOfBlocks; i++ {

			v := url.Values{}
			v.Set("id", s_n)
			i_binary := strconv.Itoa(i + 1)
			wi_bytes := []byte(i_binary)
			//fmt.Println("input:", wi_bytes)

			//h(wi) calculation
			//hash := sha256.New()
			hash := hmac.New(sha256.New, skh.Bytes())
			hash.Write(wi_bytes)
			h := pairing.NewG1().SetFromHash(hash.Sum(nil))
			//fmt.Println("h_in:", hash.Sum(nil))

			//block retreival and block prod calculation
			block := filearr[i]
			numSectors := len(block)
			prod := pairing.NewG1().Set1()
			for j := 0; j < numSectors; j++ {
				x_j := x[j]
				u_j := pairing.NewG1().PowZn(g1, x_j)
				m_ij := block[j]
				v.Add("m", m_ij.String())
				v.Add("x", x_j.String())
				//fmt.Println("m: ", m_ij)
				//fmt.Println("x:", x_j)
				//fmt.Println("u:", u_j)
				prod = pairing.NewG1().Mul(prod, pairing.NewG1().PowZn(u_j, m_ij))
			}

			tag := pairing.NewG1().PowZn(pairing.NewG1().Mul(h, prod), skt)
			//fmt.Println("tag: ", tag)
			v.Add("tag", tag.String())
			s := v.Encode()
			fmt.Println("Sending Block Number : ", i+1)
			dest := "http://127.0.0.1:8003/upload"
			conttype := "application/x-www-form-urlencoded"
			ret = functions.Connect(dest, s, conttype)
			fmt.Println(ret)

		}

		timeTrack(start, "Tag Gen Time")

	}
	x()

	z := func() {
		// Open the CSV file
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Read the CSV data into a slice of instances
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		start := time.Now()

		var datapoints [][]float64
		for _, record := range records {
			coords := make([]float64, len(record))
			for i, v := range record {
				x, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
				if err != nil {
					//log.Fatal(err)
					x = 0.0
				}
				coords[i] = x
			}
			//fmt.Println(coords)
			datapoints = append(datapoints, coords)
			//fmt.Println(datapoints)

		}
		//fmt.Println(datapoints)
		datastring, _ := functions.Float64SliceToString(datapoints)
		//fmt.Println(datastring)

		timeTrack(start, "Tree Gen")

		v := url.Values{}
		v.Set("tree", datastring)
		s := v.Encode()
		dest := "http://127.0.0.1:8003/tree"
		conttype := "application/x-www-form-urlencoded"
		ret = functions.Connect(dest, s, conttype)
		fmt.Println(ret)
	}

	z()

	if err := os.Truncate("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/noOfBlocks", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	functions.FileCrWr("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/noOfBlocks", strconv.Itoa(noOfBlocks))
	fmt.Println("File sent to Server Successfully! Total Number of Blocks = ", strconv.Itoa(noOfBlocks))

	y := func() {

		hsh := "\nconnecting to the server http://127.0.0.1:8003/hash...\n"
		hshtype := ""
		hshloc := "http://127.0.0.1:8003/hash"
		ret = functions.Connect(hshloc, hsh, hshtype)
		//fmt.Println(ret)

		h_ret, flag := pairing.NewG1().SetString(strings.Split(ret, "||")[0], 10)
		//fmt.Println(h_ret)
		if !flag {
			log.Fatalf("h_ret not set")
		}
		signature := pairing.NewG1().PowZn(h_ret, skt)
		ret = ret + "||" + signature.String()

		tx1, err := instContract.SetSign(auth, ret)
		functions.CheckError(err)
		_, err = bind.WaitMined(ctx, connection, tx1)
		functions.CheckError(err)
	}
	y()

	fmt.Println("Hash Chain committed!")
}

package main

import (
	"context"
	"crypto/sha256"
	audit "finYP/cmd/contract/abigen"
	functions "finYP/utils/func"
	kdTree "finYP/utils/kdTree"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var NumOfSectors int
var NumOfBlocks int
var SectorsInLastBlock int

var pairing *pbc.Pairing
var g1 *pbc.Element
var g2 *pbc.Element
var pk *pbc.Element
var sk *pbc.Element

var m [][]string
var tag []string
var x [][]string
var h_con []byte // storing concatenated h1,...,hn

type instance = kdTree.Instance

var tree = kdTree.NewKDTree()

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("function %s took %s", name, elapsed)
}

func ReadVariables() {
	NumOfSectors = functions.TotalSectors()
	NumOfBlocks = functions.NoofBlocks()
	SectorsInLastBlock = len(m) % NumOfSectors
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nconnection to the server http://127.0.0.1:8003/ is requested ...\n\nconnection accepted ...")
	message := "\nYou are connected to http://127.0.0.1:8003/.\n"
	w.Write([]byte(message))
}

func queryHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n --- SERVER QUERY HANDLER ---")
	ReadVariables()
	r.ParseForm()

	cap_R, _ := pairing.NewG2().SetString(r.FormValue("capr"), 10)
	//fmt.Println("\nr : ", r.FormValue("capr"))
	//fmt.Println("r cap : ", cap_R)

	//tagproof := pairing.NewG1().Set1()
	dataproof := pairing.NewGT().Set1()

	//fmt.Println("query index i : ", r.FormValue("i"))
	i_query, err := strconv.Atoi(r.FormValue("i"))
	functions.CheckError(err)

	nu, _ := pairing.NewZr().SetString(r.FormValue("nu"), 10)
	//fmt.Println("nu (\u03BD): ", r.FormValue("nu"))

	block_num := len(m) - NumOfBlocks + i_query - 1
	temp_m := m[block_num]
	temp_x := x[block_num]
	s := len(temp_m)

	start := time.Now()

	//Tag Proof Calculation
	temp_tag, flag := pairing.NewG1().SetString(tag[block_num], 10)
	//fmt.Println("temp tag: ", temp_tag)
	if !flag {
		log.Fatalf("temp tag not set")
	}
	tagproof := pairing.NewG1().PowZn(temp_tag, nu)

	//Data Proof Calculation
	for j := 0; j < s; j++ {
		mij, _ := pairing.NewZr().SetString(temp_m[j], 10)
		//fmt.Println("mij: ", mij)
		mpj := pairing.NewZr().Mul(nu, mij)
		xj, _ := pairing.NewZr().SetString(temp_x[j], 10)
		//fmt.Println("xj: ", xj)
		uj := pairing.NewG1().PowZn(g1, xj)
		//fmt.Println("uj: ", uj)
		elepair := pairing.NewGT().Pair(uj, cap_R)
		//fmt.Println("elepair: ", elepair)
		temp_dp := pairing.NewGT().PowZn(elepair, mpj)
		//fmt.Println("dataproof: ", dataproof)
		dataproof = pairing.NewGT().Mul(dataproof, temp_dp)
		//fmt.Println("dataproofAggr: ", dataproofAggr)
	}

	//fmt.Println("tag proof: ", tagproof.String())
	//fmt.Println("data proof: ", dataproof.String())

	timeTrack(start, "Proof Gen")

	fmt.Fprintf(w, tagproof.String()+"||"+dataproof.String())

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.FormValue("q")
	query, _ := functions.StringToFloat64Slice(queryString)
	min := query[0]
	max := query[1]
	searchResult := tree.RangeSearch(min, max)
	var results [][]float64
	for _, inst := range searchResult {
		coords := inst.Coords
		results = append(results, coords)
	}
	resultString, _ := functions.Float64SliceToString(results)
	fmt.Fprintf(w, resultString)
}

func treeHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	datarecstring := r.PostForm["tree"]
	var instances []*instance
	datarec, _ := functions.StringToFloat64Slice(datarecstring[0])
	for _, record := range datarec {
		coords := make([]float64, len(record))
		for i, v := range record {
			x := v
			coords[i] = x
		}
		instances = append(instances, &kdTree.Instance{Coords: coords})
	}

	//Constructing KD Tree
	for _, inst := range instances {
		tree.Insert(inst)
	}

	tree.PrintInorder()

	message := "\nKd Tree received by server\n"

	w.Write([]byte(message))

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	//x = nil

	fmt.Println("\n ---  NEW BLOCK RECEIVED  ---")

	// NOTE: Invoke ParseForm or ParseMultipartForm before reading form values
	r.ParseForm()
	/*
		Loops through r.Form object to print (id, m_id, sigma_id)
		pairs of url-encoded data. Note that these include all data
		sent through request url and request body
	*/

	var mAndtag string
	var tempm []string
	var tempx []string

	for i := 0; i < len(r.PostForm["m"]); i++ {
		//fmt.Println("m : ", r.PostForm["m"][i])

		//m = append(m, r.PostForm["m"][i])
		tempm = append(tempm, r.PostForm["m"][i])
		tempx = append(tempx, r.PostForm["x"][i])
		//fmt.Println("\ntemp m : ", tempm)
		mAndtag += r.PostForm["m"][i]
	}

	m = append(m, tempm)
	x = append(x, tempx)
	tag = append(tag, r.FormValue("tag"))
	//fmt.Println("\nm : ", m)
	//fmt.Println("tag : ", tag)

	mAndtag += r.FormValue("tag")
	hash := sha256.New()
	hash.Write([]byte(mAndtag))

	h_con = append(h_con, hash.Sum(nil)...)

	message := "\nData just sent by you received at server http://127.0.0.1:8003/upload.\n"

	w.Write([]byte(message))
}

func hashHandler(w http.ResponseWriter, r *http.Request) {
	h := pairing.NewG1().SetFromHash(h_con)
	//fmt.Println(h)
	signature := pairing.NewG1().PowZn(h, sk)
	//fmt.Println("sending : ", h.String()+"||"+signature.String())
	fmt.Fprintf(w, h.String()+"||"+signature.String())
	//fmt.Println("Hashing work done by server!")
}

func main() {
	instContract, auth, connection := ContractVars()
	ctx := context.Background()

	paramsByte, GetGVal, GetUVal, err := instContract.GetParamsGU(nil)
	functions.CheckError(err)

	pairing, err = pbc.NewPairingFromString(string(paramsByte))
	functions.CheckError(err)

	g1 = pairing.NewG1().SetBytes(GetGVal)
	g2 = pairing.NewG1().SetBytes(GetUVal)
	sk = pairing.NewZr().Rand()
	pk = pairing.NewG2().PowZn(g2, sk)

	//fmt.Println("g1 : ", g1)
	//fmt.Println("g2 : ", g2)

	//functions.FileCrWr("/home/ajay/Desktop/src/contract/tmp/private/server/skS", sk.Bytes())
	//functions.FileCrWr("/home/ajay/Desktop/src/contract/tmp/noOfBlocks", []byte(strconv.Itoa(0)))

	tx, err := instContract.SetPKS(auth, pk.Bytes())
	functions.CheckError(err)
	//fmt.Println("\nTxn hash is: ", tx.Hash())
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)
	GetPKSVal, err := instContract.GetPKS(nil)
	functions.CheckError(err)
	fmt.Println("\nGetPKS() Returned: ", GetPKSVal)
	fmt.Println("Server up and running!")

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/tree", treeHandler)
	http.HandleFunc("/hash", hashHandler)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/search", searchHandler)
	if err := http.ListenAndServe(":8003", nil); err != nil {
		log.Fatal(err)
	}
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

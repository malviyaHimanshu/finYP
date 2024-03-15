package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	audit "finYP/cmd/contract/abigen"
	functions "finYP/utils/func"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var NumOfSectors int = functions.TotalSectors()
var NumOfBlocks int
var SectorsInLastBlock int

var pairing *pbc.Pairing
var g1 *pbc.Element
var g2 *pbc.Element
var skh *pbc.Element
var random_r *pbc.Element
var pk *pbc.Element
var tagproofAggr *pbc.Element
var dataproofAggr *pbc.Element
var h_chal *pbc.Element

func ReadVariables() {
	NumOfSectors = functions.TotalSectors()
	NumOfBlocks = functions.NoofBlocks()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("function %s took %s", name, elapsed)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

type SharedParams struct {
	Params    string // group parameter
	G         []byte // shared G
	U         []byte
	PK        []byte
	AggrSigma []byte
	AggrMU    []byte
	//structVars Shared
	I  [][]byte
	NU [][]byte
}

var Icopy [][]byte
var NUcopy [][]byte

func ContractVars() (*audit.Audit, *bind.TransactOpts, *ethclient.Client) {
	gethPath, nodeKey := functions.GethPathAndKey()

	addrByte, err := ioutil.ReadFile("/home/ajay/Desktop/src/contract/tmp/contractInfo/contract-address")
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

func verification(tagproof *pbc.Element, dataproof *pbc.Element) {

	start := time.Now()

	lhs := pairing.NewGT().Mul(dataproof, pairing.NewGT().Pair(h_chal, pk))
	rhs := pairing.NewGT().Pair(tagproof, pairing.NewG2().PowZn(g2, random_r))

	//fmt.Println("LHS = ", lhs)
	//fmt.Println("RHS = ", rhs)

	if lhs.Equals(rhs) {
		fmt.Println("AUDIT SUCCESSFUL !!")
	} else {
		fmt.Println("AUDIT FAILED !!")
	}

	timeTrack(start, "Verification")
}

func search() {

	start := time.Now()
	v := url.Values{}

	query := [][]float64{{30, 30}, {90, 60}}
	//query := [][]float64{{0, 0}, {5, 5}}
	queryString, _ := functions.Float64SliceToString(query)

	v.Set("q", queryString)
	s := v.Encode()
	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://127.0.0.1:8003/search", strings.NewReader(s))
	functions.CheckError(err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(s)))
	resp, err := client.Do(r)
	functions.CheckError(err)
	defer resp.Body.Close()

	var strmsg string
	if resp.StatusCode == http.StatusOK {
		message, err := ioutil.ReadAll(resp.Body)
		functions.CheckError(err)
		strmsg = string(message)
	}

	results, _ := functions.StringToFloat64Slice(strmsg)
	fmt.Println("Found", len(results), "nodes within search range:")
	timeTrack(start, "Search ")
	for _, inst := range results {
		fmt.Println(inst)
	}
}

func queryServer() {
	noOfBlocks := functions.NoofBlocks()

	// Let the query be (i,nu)
	// Generate i randomly between [1,noOfBlocks]

	v := url.Values{}

	mrand := random(1, noOfBlocks+1)
	i := []byte(strconv.Itoa(mrand))
	nu := pairing.NewZr().Rand()

	i_str := strconv.Itoa(mrand)
	wi_bytes := []byte(i_str)
	//fmt.Println("input:", wi_bytes)

	//h(wi) calculation
	//hash := sha256.New()
	hash := hmac.New(sha256.New, skh.Bytes())
	hash.Write(wi_bytes)
	h_in := pairing.NewG1().SetFromHash(hash.Sum(nil))
	//fmt.Println("h_in:", hash.Sum(nil))
	prod := pairing.NewZr().Mul(random_r, nu)
	h := pairing.NewG1().PowZn(h_in, prod)
	h_chal = pairing.NewG1().Mul(h_chal, h)
	//fmt.Println("h_chal:", h_chal)

	v.Set("i", string(i))
	v.Add("nu", nu.String())
	//fmt.Println("i:", string(i))
	//fmt.Println("nu:", nu)

	Icopy = append(Icopy, i)
	NUcopy = append(NUcopy, nu.Bytes())

	cap_r := pairing.NewG2().PowZn(pk, random_r)
	v.Add("capr", cap_r.String())
	s := v.Encode()
	//fmt.Println("r:", cap_r)

	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://127.0.0.1:8003/query", strings.NewReader(s))
	functions.CheckError(err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(s)))

	resp, err := client.Do(r)
	functions.CheckError(err)
	defer resp.Body.Close()

	var strmsg string
	if resp.StatusCode == http.StatusOK {
		message, err := ioutil.ReadAll(resp.Body)
		functions.CheckError(err)
		strmsg = string(message)
	}
	functions.CheckError(err)

	tagproof, _ := pairing.NewG1().SetString(strings.Split(strmsg, "||")[0], 10)

	dataproof, _ := pairing.NewGT().SetString(strings.Split(strmsg, "||")[1], 10)

	tagproofAggr = pairing.NewG1().Mul(tagproofAggr, tagproof)
	dataproofAggr = pairing.NewGT().Mul(dataproofAggr, dataproof)

	//fmt.Println("tag proof: ", tagproofAggr.String())
	//fmt.Println("data proof: ", dataproofAggr.String())

}

func main() {
	instContract, auth, connection := ContractVars()

	paramsByte, GetGVal, GetUVal, err := instContract.GetParamsGU(nil)
	functions.CheckError(err)

	GetPKOVal, err := instContract.GetPKT(nil)
	GetSKHVal, err := instContract.GetSKH(nil)
	functions.CheckError(err)

	pairing, err = pbc.NewPairingFromString(string(paramsByte))
	functions.CheckError(err)

	g1 = pairing.NewG1().SetBytes(GetGVal)
	g2 = pairing.NewG2().SetBytes(GetUVal)
	pk = pairing.NewG2().SetBytes(GetPKOVal)
	skh = pairing.NewZr().SetBytes(GetSKHVal)

	//fmt.Println("g1 : ", g1)
	//fmt.Println("g2 : ", g2)
	//fmt.Println("skh : ", skh.Bytes())

	fmt.Println("\nGetPKO() Returned: ", GetPKOVal)

	random_r = pairing.NewZr().Rand()
	h_chal = pairing.NewG1().Set1()
	tagproofAggr = pairing.NewG1().Set1()
	dataproofAggr = pairing.NewGT().Set1()

	rand.Seed(time.Now().Unix())

	for j := 0; j < 1; j++ {
		queryServer()
	}

	verification(tagproofAggr, dataproofAggr)

	search()

	structParamsVar := SharedParams{
		Params:    string(paramsByte),
		G:         GetGVal,
		U:         GetUVal,
		PK:        GetPKOVal,
		AggrSigma: tagproofAggr.Bytes(),
		AggrMU:    dataproofAggr.Bytes(),
		I:         Icopy,
		NU:        NUcopy,
	}

	byteShared, err := json.Marshal(structParamsVar)
	functions.CheckError(err)

	ctx := context.Background()

	// Send audit query
	tx, err := instContract.SendAudit(auth, []byte(byteShared))
	functions.CheckError(err)
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)

}

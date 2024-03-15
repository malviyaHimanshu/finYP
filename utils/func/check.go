package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const MaxSectors = 1000

/* Universal Functions. */

func TotalSectors() int {
	return MaxSectors
}

// function to check error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// function to create directory if not exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		CheckError(err)
	}
}

// function to create tmp folder and subsequent inner directories
func Createtmp() {
	CreateDirIfNotExist("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp")
	CreateDirIfNotExist("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/contractInfo")
	CreateDirIfNotExist("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/private")
	CreateDirIfNotExist("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/private/owner")
	CreateDirIfNotExist("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/private/server")
	CreateDirIfNotExist("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/private/auditor")
}

// function to create file 's' and write 'b' bytes in it
func FileCrWr(s string, b string) {
	f, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE, 0755)
	CheckError(err)
	_, err = f.WriteString(b)
	CheckError(err)
	defer f.Close()
}

// function to read variables from file (residing in folder tmp)
func NoofBlocks() int {
	/*paramByte, err := ioutil.ReadFile("./tmp/pbcInfo/params")
	CheckError(err)
	pairing, _ := pbc.NewPairingFromString(string(paramByte))

	gByte, err := ioutil.ReadFile("./tmp/public/gO")
	CheckError(err)
	g := pairing.NewG1().SetBytes(gByte)

	uByte, err := ioutil.ReadFile("./tmp/public/u")
	CheckError(err)
	u := pairing.NewG1().SetBytes(uByte)

	pkByte, err := ioutil.ReadFile("./tmp/public/pkO")
	CheckError(err)
	pk := pairing.NewG1().SetBytes(pkByte)*/

	noOfBlocksByte, err := ioutil.ReadFile("/Users/himanshumalviya/Desktop/finYP/cmd/contract/tmp/noOfBlocks")
	CheckError(err)
	s_noOfBlocks := string(noOfBlocksByte)
	noOfBlocks, err := strconv.Atoi(s_noOfBlocks)
	CheckError(err)

	return noOfBlocks
}

// function to connect and post new request using http
func Connect(dest string, content string, contentType string) string {
	client := &http.Client{}
	r, err := http.NewRequest("POST", dest, strings.NewReader(content))
	CheckError(err)
	r.Header.Add("Content-Type", contentType)
	r.Header.Add("Content-Length", strconv.Itoa(len(content)))

	resp, err := client.Do(r)
	CheckError(err)
	defer resp.Body.Close()

	var message []byte
	if resp.StatusCode == http.StatusOK {
		message, err = ioutil.ReadAll(resp.Body)
		CheckError(err)
	}
	return string(message)
}

// function to get path of Geth node and keystore key
func GethPathAndKey() (string, string) {
	dir, err := os.Getwd()
	CheckError(err)

	pathParent := path.Dir(dir)
	pathParent += "/node-data/node2/"
	gethPath := pathParent + "geth.ipc"

	//var fileName string
	var filePath string
	keyStorePath := pathParent + "keystore/"

	err = filepath.Walk(keyStorePath, func(path string, info os.FileInfo, err error) error {
		CheckError(err)
		if !info.IsDir() {
			fmt.Printf("visited file or dir: %q\n", path)
			//fileName = info.Name()
			filePath = path
		}
		return nil
	})
	CheckError(err)

	var key []byte
	keyRead := func() {
		var err error
		key, err = ioutil.ReadFile(filePath)
		CheckError(err)
	}
	keyRead()
	return gethPath, string(key)
}

// function to check timings
/*func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("function %s took %s", name, elapsed)
}*/

func Float64SliceToString(f [][]float64) (string, error) {
	b, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func StringToFloat64Slice(s string) ([][]float64, error) {
	var f [][]float64
	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

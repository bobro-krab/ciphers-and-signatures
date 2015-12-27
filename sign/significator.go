package sign

import (
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"os"
	"zi/shifr"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Significator interface {
	GenSign(hash int) []int
	CheckSign(sign []int, fileHash int) bool // is signature valid

	FileType() string
	Init()
	Key() []byte
	LoadKey([]byte)
}

func Checksum(file []byte) int {
	return int(crc32.ChecksumIEEE(file))
}

func SignupFile(filename string, s Significator) {
	fmt.Println("Signup file")
	signatureFilename := filename + "." + s.FileType() + ".sign"
	s.Init()

	fileToSign := shifr.GetBytesFromFile(filename)
	hash := Checksum(fileToSign)
	signature := s.GenSign(hash)
	fmt.Println(filename, "hash is:", hash)
	// fmt.Println(filename, "sign is", signature)
	signatureFile, _ := os.Create(signatureFilename)
	defer signatureFile.Close()

	// // save signature to separate file
	// writer := bufio.NewWriter(signatureFile)
	// for _, v := range signature {
	// 	writer.WriteString(strconv.Itoa(v))
	// }
	// writer.WriteString("hello")

	enc := gob.NewEncoder(signatureFile)
	enc.Encode(signature)

	// save key to file
	keyFile, err := os.Create(filename + "." + s.FileType() + ".key")
	check(err)
	defer keyFile.Close()
	keyFile.Write(s.Key())
}

func CheckupSignature(filename string, s Significator) bool {
	// read signature from file
	signatureFilename := filename + "." + s.FileType() + ".sign"
	signFile, _ := os.Open(signatureFilename)
	defer signFile.Close()
	signature := make([]int, 4)
	dec := gob.NewDecoder(signFile)
	dec.Decode(&signature)

	// for _, v := range signature {
	// 	fmt.Println("some ", v)
	// }

	fileBytes := shifr.GetBytesFromFile(filename)

	keyFile, err := os.Open(filename + "." + s.FileType() + ".key")
	check(err)
	defer keyFile.Close()

	// read key, and set our setting to that
	stat, _ := keyFile.Stat()
	key := make([]byte, stat.Size())
	keyFile.Read(key)
	s.LoadKey(key)

	fmt.Println(filename, "checkup hash is:", Checksum(fileBytes))
	// CHECK FOR EQUALITY
	result := s.CheckSign(signature, Checksum(fileBytes))
	fmt.Println("result is ", result)
	return result
}

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
	Init()
	GenSign(hash int) []int
	CheckSign(sign []int, fileHash int) bool // is signature valid
	FileType() string

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
	fmt.Println(filename, "hash is:", hash)
	fmt.Println(filename, "sign is", s.GenSign(hash))
	signatureFile, _ := os.Create(signatureFilename)
	defer signatureFile.Close()

	keyFile, err := os.Create(filename + "." + s.FileType() + ".key")
	check(err)
	defer keyFile.Close()

	// save signature to separate file
	keyFile.Write(s.Key())
	encoder := gob.NewEncoder(signatureFile)
	encoder.Encode(s.GenSign(hash))
}

func CheckupSignature(filename string, s Significator) bool {
	signatureFilename := filename + "." + s.FileType() + ".sign"
	signFile, _ := os.Open(signatureFilename)
	defer signFile.Close()

	decoder := gob.NewDecoder(signFile)
	signature := []int{}
	decoder.Decode(&signature)
	for i := range signature {
		fmt.Println("i:", i, "value:", signature[i])
	}

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

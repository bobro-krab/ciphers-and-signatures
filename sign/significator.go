package sign

import (
	"fmt"
	"os"
	"strconv"
	"zi/shifr"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Significator interface {
	Init()
	Checksum(file []byte) int
	GenSign(hash int) []int
	GetHashFromSign(sign int) string
	FileType() string

	Key() []byte
	LoadKey([]byte)
}

func SignupFile(filename string, s Significator) {
	signatureFilename := filename + "." + s.FileType() + ".sign"
	s.Init()
	fileToSign := shifr.GetBytesFromFile(filename)
	hash := s.Checksum(fileToSign)
	fmt.Println(filename, "hash is:", hash)
	fmt.Println(filename, "sign is", s.GenSign(hash))
	signatureFile, _ := os.Create(signatureFilename)
	defer signatureFile.Close()

	keyFile, err := os.Create(filename + "." + s.FileType() + ".key")
	check(err)
	defer keyFile.Close()

	keyFile.Write(s.Key())
	signatureFile.WriteString(string(s.GenSign(hash)[0]))
}

func CheckSign(filename string, s Significator) bool {
	signatureFilename := filename + "." + s.FileType() + ".sign"
	signFile, _ := os.Open(signatureFilename)
	defer signFile.Close()

	stat, _ := signFile.Stat()
	signBytes := make([]byte, stat.Size())
	signFile.Read(signBytes)
	sign := string(signBytes)

	fileBytes := shifr.GetBytesFromFile(filename)

	fmt.Println(filename, "read sign is'", sign, "'")
	temp, _ := strconv.Atoi(sign)
	fmt.Println("temp is ", temp)

	keyFile, err := os.Open(filename + "." + s.FileType() + ".key")
	check(err)
	defer keyFile.Close()

	// read key, and set our setting to that
	stat, _ = keyFile.Stat()
	key := make([]byte, stat.Size())
	keyFile.Read(key)
	s.LoadKey(key)

	hash1 := s.GetHashFromSign(int(temp))
	hash2 := strconv.Itoa(s.Checksum(fileBytes))
	fmt.Println("hash1 and hash2 is", hash1, hash2)
	if hash1 == hash2 {
		fmt.Println("Signature is right")
		return true
	} else {
		fmt.Println("Signature is NOT right")
		return false
	}

	return true
}

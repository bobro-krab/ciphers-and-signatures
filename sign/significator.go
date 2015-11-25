package sign

import (
//"fmt"
//"zi/crypto"
)

type Significator interface {
	Init()
	Checksum(file []byte) int
}

func SignupFile(filename string, s Significator) {
	fileToSign := getBytesFromFile(filename)
	hash := s.Checksum(fileToSign)
}

func CheckSign() bool {
	return true
}

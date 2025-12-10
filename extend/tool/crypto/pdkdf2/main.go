package main

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	fmt.Printf("passwd == %s \n", EncodePassword("lewaimai123.", "a9kwR9TlJB"))
}

func EncodePassword(password, salt string) string {
	newPasswd := pbkdf2.Key([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", newPasswd)
}

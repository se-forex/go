package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
)

func genSalt(n int) string {
	b := make([]byte, n)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}

func genHashPass(p, s []byte) []byte {
	passHash := md5.New()
	passHash.Write(md5.New().Sum([]byte(p)))
	passHash.Write(sha1.New().Sum([]byte(s)))

	return passHash.Sum(nil)
}

func main() {

	pass := "Se2l0f6"
	salt := genSalt(10)

	newPass := "Se2l0f61"

	// io.WriteString(h, "love")
	// st1 := h.Sum([]byte("love"))
	// st2 := h.Sum([]byte("123123"))

	// if bytes.Equal(st1, st2) {
	// 	fmt.Println("!!!!")
	// }

	// fmt.Println(pass)
	// fmt.Printf("%x \n", md5.New().Sum([]byte(pass)))

	// fmt.Println(salt)
	// fmt.Printf("%x \n", sha1.New().Sum([]byte(salt)))

	hp1 := genHashPass(md5.New().Sum([]byte(pass)), sha1.New().Sum([]byte(salt)))
	hp2 := genHashPass(md5.New().Sum([]byte(newPass)), sha1.New().Sum([]byte(salt)))

	fmt.Println("Initial passHash:")
	fmt.Printf("%x \n", hp1)
	fmt.Println("New passHash:")
	fmt.Printf("%x \n", hp2)

	if bytes.Equal(hp1, hp2) {
		fmt.Println("\nPasswors equal")
	} else {
		fmt.Println("\nPasswors not equal")
	}
}

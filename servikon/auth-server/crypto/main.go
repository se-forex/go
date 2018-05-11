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

	sss := "a"

	pass := "123123"
	salt := genSalt(10)

	serverSalt := "MIfpaOS&UE02j30*YdfQA(PE8fQP(#&@Uhsal"

	newPass := "123123"

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

	hp1 := genHashPass(genHashPass(md5.New().Sum([]byte(pass)), sha1.New().Sum([]byte(salt))), md5.New().Sum([]byte(serverSalt)))
	hp2 := genHashPass(genHashPass(md5.New().Sum([]byte(newPass)), sha1.New().Sum([]byte(salt))), md5.New().Sum([]byte(serverSalt)))

	fmt.Println("Initial passHash:")
	fmt.Printf("%x \n", hp1)
	// fmt.Print(len(hp1), " ")
	// fmt.Print(len(sss), " ")
	// fmt.Println(len(string(hp1) + sss))
	// fmt.Printf("%x \n", string(hp1))
	fmt.Println("New passHash:")
	fmt.Printf("%x \n", hp2)

	fmt.Print("len h2: ")
	fmt.Printf("%x \n", len(string(hp2)))
	fmt.Println("len sss:", len(sss))
	fmt.Printf("%x \n", string(hp2)+sss)

	nPass := string(hp2) + sss
	fmt.Println("sss")
	fmt.Printf("%x \n", nPass[:(len(nPass)-len(sss))])

	if bytes.Equal(hp1, hp2) {
		fmt.Println("\nPasswors equal")
	} else {
		fmt.Println("\nPasswors not equal")
	}
}

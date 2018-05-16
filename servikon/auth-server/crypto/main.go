package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

const serverSalt = "MIfpaOS&UE02j30*YdfQA(PE8fQP(#&@Uhsal"

func genSalt(n int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	r.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}

func genHashPass(p, s []byte) []byte {
	passHash := md5.New()
	passHash.Write(md5.New().Sum([]byte(p)))
	passHash.Write(sha1.New().Sum([]byte(s)))

	return passHash.Sum(nil)
}

//add email string
func procPassForDB(p []byte) string {
	sss := "a"
	return string(p) + sss
}

func procPassAfterDB(p string) []byte {
	sss := "a"
	return []byte(p[:(len(p) - len(sss))])
}

func main() {

	pass := "123123"
	salt := genSalt(10)

	for i := 0; i < 10; i++ {
		fmt.Printf("%x \n", genSalt(10))
	}

	newPass := "111111"

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

	fmt.Println("Proc pass for DB:")
	fmt.Printf("%x \n", procPassForDB(hp1))

	fmt.Println("Proc pass after DB:")
	fmt.Printf("%x \n", procPassAfterDB(procPassForDB(hp1)))

	if bytes.Equal(hp1, hp2) {
		fmt.Println("\nPasswors equal")
	} else {
		fmt.Println("\nPasswors not equal")
	}
}

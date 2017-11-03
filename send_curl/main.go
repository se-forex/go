package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
const prod = "http://srv-osb12c-vip:8000"
const pprod = "http://srv-pposb12-balancerHA:8000"
const so = "http://srv-so1osb12-01:7004"

type tmpJSON struct {
	Name string `json:"name"`
	Val  string `json:"env"`
}

type Request struct {
	Name string    `json:"name"`
	Val  InRequest `json:"val"`
}

type InRequest struct {
	Prod  string `json:"prod"`
	Pprod string `json:"pprod"`
	So    string `json:"so"`
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func formJSON() {
	tJSON := tmpJSON{}
	fJSON := Request{}

	file, err := os.Open("./send_curl/Proxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		json.Unmarshal([]byte(scanner.Text()), &tJSON)
		fJSON.Name = tJSON.Name
		fJSON.Val.Prod = prod + tJSON.Val
		fJSON.Val.Pprod = pprod + tJSON.Val
		fJSON.Val.So = so + tJSON.Val
		sendRequest(fJSON)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sendRequest(resJSON Request) {
	idVal := randString(33)
	jsonStr := []byte(`{"name":"` + resJSON.Name + `","val":{"prod":"` + resJSON.Val.Prod + `","pprod":"` + resJSON.Val.Pprod + `","so":"` + resJSON.Val.So + `"}}`)
	req, err := http.NewRequest("PUT", "http://127.0.0.1:8080/env/"+string(idVal), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	fmt.Println("Name:", string(idVal))
	fmt.Println("JSON:", string(jsonStr))
}

func main() {
	formJSON()
}

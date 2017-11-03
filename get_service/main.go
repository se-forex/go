package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type InResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []InResponse `json:"rows"`
}

func daemonize() {
	// check command line flags, configuration etc.

	// short delay to avoid race condition between os.StartProcess and os.Exit
	// can be omitted if the work done above amounts to a sufficient delay
	time.Sleep(1 * time.Second)

	if os.Getppid() != 1 {
		// I am the parent, spawn child to run as daemon
		binary, err := exec.LookPath(os.Args[0])
		if err != nil {
			log.Fatalln("Failed to lookup binary:", err)
		}
		_, err = os.StartProcess(binary, os.Args, &os.ProcAttr{Dir: "", Env: nil,
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, Sys: nil})
		if err != nil {
			log.Fatalln("Failed to start process:", err)
		}
		os.Exit(0)
	} else {
		// I am the child, i.e. the daemon, start new session and detach from terminal
		_, err := syscall.Setsid()
		if err != nil {
			log.Fatalln("Failed to create new session:", err)
		}
		file, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if err != nil {
			log.Fatalln("Failed to open /dev/null:", err)
		}
		syscall.Dup2(int(file.Fd()), int(os.Stdin.Fd()))
		syscall.Dup2(int(file.Fd()), int(os.Stdout.Fd()))
		syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd()))
		file.Close()
	}
}

func findValue(env, srv string) string {
	fullJSON := new(Response)
	var resEnv string
	res, err := http.Get("http://localhost:5984/env/_design/" + env + "/_view/values/")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &fullJSON)

	for i := 0; i < fullJSON.TotalRows; i++ {
		if fullJSON.Rows[i].Name == srv {
			resEnv = fullJSON.Rows[i].Value
		}
	}

	return resEnv
}

func waitHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, findValue(r.URL.Query().Get("env"), r.URL.Query().Get("srv")))
}

func main() {
	daemonize()

	http.HandleFunc("/", waitHTTP)
	http.ListenAndServe(":5555", nil)
}

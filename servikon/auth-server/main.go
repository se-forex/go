package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

//Global DB object
var db *sql.DB

type User struct {
	Uid      string
	Username string
	Email    string
	Pass     string
}

//DB section
// 1. DB init
func initDb() {
	var err error

	db, err = sql.Open("postgres", "postgres://user:user@localhost:1521/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB successfully connected!")
}

// Http server section
func getUserByEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side

	user := new(User)
	var out []byte
	var err error

	for k, v := range r.Form {
		if k == "email" {
			*user, err = searchUserByEmail(strings.Join(v, ""))
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	out, err = json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprintf(w, string(out))

	//fmt.Fprintf(w, "%s, %s, %s, %s\n", user.Uid, user.Username, user.Email, user.Pass) // send data to client side
}

func searchUserByEmail(email string) (User, error) {
	user := new(User)

	// Read data from DB
	rows, errQr := db.Query("SELECT uid, username, email, pass FROM records.users WHERE email='" + email + "'")
	if errQr != nil {
		log.Fatal(errQr)
		return *user, errQr
	}
	defer rows.Close()

	for rows.Next() {
		errSt := rows.Scan(&user.Uid, &user.Username, &user.Email, &user.Pass)
		if errSt != nil {
			log.Fatal(errSt)
			return *user, errSt
		}
	}

	return *user, nil
}

func main() {
	// Create DB connection
	initDb()
	defer db.Close()

	// start Http server
	http.HandleFunc("/get", getUserByEmail)  // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

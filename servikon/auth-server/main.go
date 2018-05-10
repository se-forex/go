package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	Uid      string
	Username string
	Email    string
	Pass     string
}

type Users struct {
	Id []User
}

func (st *Users) AddUser(user User) {
	st.Id = append(st.Id, user)
}

func (st *Users) GetUserFromEmail(email string) (User, error) {
	for i := range st.Id {
		if st.Id[i].Email == email {
			return st.Id[i], nil
		}
	}
	return st.Id[0], nil
}

func getUserByEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side

	user := new(User)
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
	fmt.Fprintf(w, "%s, %s, %s, %s\n", user.Uid, user.Username, user.Email, user.Pass) // send data to client side
}

func searchUserByEmail(email string) (User, error) {
	user := new(User)
	db, errDb := sql.Open("postgres", "postgres://user:user@localhost:1521/postgres?sslmode=disable")
	if errDb != nil {
		log.Fatal(errDb)
		return *user, errDb
	}

	// Read data fron DB
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
	// start Http server
	http.HandleFunc("/get", getUserByEmail)  // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// Users init
	// var users = Users{}

	// PG DB init
	// db, err := sql.Open("postgres", "postgres://user:user@localhost:1521/postgres?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Read data fron DB
	// rows, err := db.Query("SELECT uid, username, email, pass FROM records.users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	user := new(User)
	// 	err := rows.Scan(&user.Uid, &user.Username, &user.Email, &user.Pass)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	users.AddUser(*user)
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, user := range users.Id {
	// 	fmt.Printf("%s, %s, %s, %s\n", user.Uid, user.Username, user.Email, user.Pass)
	// }

}

package main

import (
	"database/sql"
	"fmt"
	"log"

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

func main() {
	// Users init
	var users = Users{}

	// PG DB init
	db, err := sql.Open("postgres", "postgres://user:user@localhost:1521/users?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Read data fron DB
	rows, err := db.Query("SELECT uid, username, email, pass FROM records.users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Uid, &user.Username, &user.Email, &user.Pass)
		if err != nil {
			log.Fatal(err)
		}

		users.AddUser(*user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, user := range users.Id {
		fmt.Printf("%s, %s, %s, %s\n", user.Uid, user.Username, user.Email, user.Pass)
	}

}

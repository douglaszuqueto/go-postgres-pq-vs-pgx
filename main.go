package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/douglaszuqueto/go-postgres-pq-vs-pgx-pq-vs-pgx/pkg/storage"
)

var (
	dbPq = storage.NewGoPq()
	dbPg = storage.NewGoPg()
)

func main() {
	fmt.Println("GoPg")

	// testPq()
	// testPg()

	http.HandleFunc("/pq", handlerPq)
	http.HandleFunc("/pg", handlerPg)

	http.ListenAndServe(":3000", nil)
}

func handlerPg(w http.ResponseWriter, r *http.Request) {
	user, err := dbPg.GetUsers()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(user)
}

func handlerPq(w http.ResponseWriter, r *http.Request) {
	user, err := dbPq.GetUsers()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(user)
}

var size = 10

func testPq() {
	db := storage.NewGoPq()

	for i := 0; i < size; i++ {
		user, err := db.GetUser("2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
		if err != nil {
			panic(err)
		}

		fmt.Println(user.ID, user.Username)
	}

	for i := 0; i < size; i++ {
		user, err := db.GetUsers()
		if err != nil {
			panic(err)
		}

		fmt.Println(len(user))
	}
}

func testPg() {
	db := storage.NewGoPg()

	for i := 0; i < size; i++ {
		user, err := db.GetUser("2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
		if err != nil {
			panic(err)
		}

		fmt.Println(user.ID, user.Username)
	}

	for i := 0; i < size; i++ {
		user, err := db.GetUsers()
		if err != nil {
			panic(err)
		}

		fmt.Println(len(user))
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/douglaszuqueto/go-postgres-pq-vs-pgx/pkg/storage"
)

var (
	dbPq = storage.NewGoPq()
	dbPg = storage.NewGoPg()
)

func main() {
	fmt.Println("Go PostgreSQL test drivers")

	testPq()
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

	// for i := 0; i < size; i++ {
	// 	user, err := db.GetUser("2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(user.ID, user.Username)
	// }

	// for i := 0; i < size; i++ {
	// 	user, err := db.GetUsers()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(len(user))
	// }

	getUser2(db, 1)
	getUser2(db, 10)
	getUser2(db, 15)
}

func testPg() {
	db := storage.NewGoPg()

	// for i := 0; i < size; i++ {
	// 	ts := time.Millisecond * 8
	// 	ctx, cancel := context.WithTimeout(context.Background(), ts)
	// 	defer cancel()

	// 	user, err := db.GetUser(ctx, "2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	fmt.Println(i, user.ID, user.Username)
	// }

	getUser(db, 5)
	getUser(db, 10)
	getUser(db, 15)

	// for i := 0; i < size; i++ {
	// 	user, err := db.GetUsers()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(len(user))
	// }
}

func getUser(db *storage.GoPg, n int) {
	ts := time.Millisecond * time.Duration(n)

	ctx, cancel := context.WithTimeout(context.Background(), ts)
	defer cancel()

	user, err := db.GetUser(ctx, "2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user.ID, user.Username, ts)
}

func getUser2(db *storage.GoPq, n int) {
	ts := time.Millisecond * time.Duration(n)

	ctx, cancel := context.WithTimeout(context.Background(), ts)
	defer cancel()

	user, err := db.GetUser(ctx, "2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user.ID, user.Username, ts)
}

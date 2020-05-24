package main_test

import (
	"testing"

	"github.com/douglaszuqueto/go-postgres-pq-vs-pgx/pkg/storage"
)

var size = 10

func BenchmarkPqGetUser(b *testing.B) {
	db := storage.NewGoPq()

	for n := 0; n < b.N; n++ {
		_, err := db.GetUser("2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkPqGetUsers(b *testing.B) {
	db := storage.NewGoPq()

	for n := 0; n < b.N; n++ {
		_, err := db.GetUsers()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkPgGetUser(b *testing.B) {
	db := storage.NewGoPq()

	for n := 0; n < b.N; n++ {
		_, err := db.GetUser("2e4418d4-0deb-4131-a9f6-d173c15d8c3b")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkPgGetUsers(b *testing.B) {
	db := storage.NewGoPg()

	for n := 0; n < b.N; n++ {
		_, err := db.GetUsers()
		if err != nil {
			panic(err)
		}
	}
}

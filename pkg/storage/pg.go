package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4"
)

// GoPg GoPg
type GoPg struct {
	db *pgx.Conn
}

// NewGoPg NewGoPg
func NewGoPg() *GoPg {
	db := &GoPg{}
	db.connect()

	return db
}

// GetUser GetUser
func (s *GoPg) GetUser(ctx context.Context, id string) (User, error) {
	var u User

	query := `
		SELECT
			id,
			username,
			email
		FROM
			public.user
		WHERE
			id = $1
	`

	err := s.doQueryRow("[PG] GetUser", ctx, query, id).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
	)

	if err != nil {
		return u, HandlePSQLError(err)
	}

	return u, nil
}

// GetUsers GetUsers
func (s *GoPg) GetUsers() ([]User, error) {
	var l []User

	query := `
		SELECT
			id,
			username,
			email
		FROM
			public.user
	`

	rows, err := s.doQuery("[PG] GetUsers", query)
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
		)

		if err != nil {
			panic(err)
		}

		l = append(l, u)
	}

	if err != nil {
		return l, HandlePSQLError(err)
	}

	return l, nil
}

func (s *GoPg) doQuery(name, query string, args ...interface{}) (pgx.Rows, error) {
	start := time.Now()
	rows, err := s.db.Query(context.Background(), query, args...)
	logQuery(name, time.Since(start), args...)

	return rows, err
}

func (s *GoPg) doQueryRow(name string, ctx context.Context, query string, args ...interface{}) pgx.Row {
	start := time.Now()
	row := s.db.QueryRow(ctx, query, args...)
	logQuery(name, time.Since(start), args...)

	return row
}

func (s *GoPg) connect() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_DATABASE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbDatabase,
	)

	var err error
	s.db, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalln(err)
	}

	// s.db.SetMaxIdleConns(5)
	// s.db.SetMaxOpenConns(5)
	// s.db.SetConnMaxLifetime(5 * time.Minute)

	err = s.db.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}

package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// GoPq GoPq
type GoPq struct {
	db *sql.DB
}

// NewGoPq NewGoPq
func NewGoPq() *GoPq {
	db := &GoPq{}
	db.connect()

	return db
}

// GetUser GetUser
func (s *GoPq) GetUser(ctx context.Context, id string) (User, error) {
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

	err := s.doQueryRow(ctx, "[PQ] GetUser", query, id).Scan(
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
func (s *GoPq) GetUsers() ([]User, error) {
	var u User
	var l []User

	query := `
		SELECT
			id,
			username,
			email
		FROM
			public.user
	`

	rows, err := s.doQuery("[PQ] GetUsers", query)
	defer rows.Close()

	for rows.Next() {
		_ = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
		)

		l = append(l, u)
	}

	if err != nil {
		return l, HandlePSQLError(err)
	}

	return l, nil
}

func (s *GoPq) doQuery(name, query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := s.db.Query(query, args...)
	logQuery(name, time.Since(start), args...)

	return rows, err
}

func (s *GoPq) doQueryRow(ctx context.Context, name, query string, args ...interface{}) *sql.Row {
	start := time.Now()
	row := s.db.QueryRowContext(ctx, query, args...)
	logQuery(name, time.Since(start), args...)

	return row
}

func (s *GoPq) connect() {
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
	s.db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	s.db.SetMaxIdleConns(5)
	s.db.SetMaxOpenConns(5)
	s.db.SetConnMaxLifetime(5 * time.Minute)

	err = s.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

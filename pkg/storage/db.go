package storage

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

var (
	// ErrAlreadyExists ErrAlreadyExists
	ErrAlreadyExists = errors.New("object already exists")

	// ErrDoesNotExist ErrDoesNotExist
	ErrDoesNotExist = errors.New("object does not exist")

	// ErrQueryIsCanceled ErrQueryIsCanceled
	ErrQueryIsCanceled = errors.New("query is canceled")
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel)
}

func logQuery(query string, duration time.Duration, args ...interface{}) {
	log.WithFields(log.Fields{
		"query":    query,
		"duration": duration,
	}).Debug("sql query executed")
}

// HandlePSQLError handlePSQLError
// https://github.com/brocaar/lora-app-server/blob/master/internal/storage/errors.go#L41
func HandlePSQLError(err error) error {
	if err == sql.ErrNoRows {
		return ErrDoesNotExist
	}

	switch err := err.(type) {
	case *pq.Error:
		switch err.Code.Name() {
		case "unique_violation":
			return errors.Wrap(ErrAlreadyExists, err.Constraint)
		case "foreign_key_violation":
			return ErrDoesNotExist
		case "query_canceled":
			return ErrQueryIsCanceled
		default:
			log.Println(err.Code.Name())
			return err
		}
	}

	return err
}

package posts_creation

import (
	"golang_boilerplate/logger"

	"github.com/jmoiron/sqlx"
	"golang_boilerplate/errors"
	"github.com/newrelic/go-agent"
	"golang_boilerplate/instrumentation"
)

type PostsRepository interface {
	Create(*Post, newrelic.Transaction) (string, error)
}

type PostsRepositoryImpl struct {
	db     *sqlx.DB
	logger logger.Log
}

func NewPostsRepository(db *sqlx.DB, logger logger.Log) PostsRepository {
	return &PostsRepositoryImpl{db: db, logger: logger}
}

const createQuery = "INSERT into posts (topic, body) values ($1, $2) returning id"

func (tr *PostsRepositoryImpl) Create(post *Post, txn newrelic.Transaction) (string, error) {
	defer instrumentation.StartDataSegmentNowForPostgres("insert", "posts", txn).End()

	tx, err := tr.db.Begin()
	if err != nil {
		tr.logger.Errorf("Failed to create txn for creating post: %v", err)
		return "", errors.NewError(errors.GenericServiceError, err)
	}

	id := ""
	row := tx.QueryRow(createQuery,
		post.Topic,
		post.Body)
	err = row.Scan(&id)

	if err != nil {
		tr.logger.Warn("Rolling back txn for creating post: %v", err)
		tx.Rollback()
		return "", errors.NewError(errors.GenericServiceError, err)
	}

	err = tx.Commit()
	if err != nil {
		tr.logger.Errorf("Failed to commit txn for creating post: %v", err)
		err = tx.Rollback()
		if err != nil {
			return "", errors.NewError(errors.GenericServiceError, err)
		}
	} else {
		tr.logger.Debug("Successfully created Post in DB")
	}

	return id, nil
}

package posts_creation

import (
	"testing"
	_ "database/sql"
	"github.com/stretchr/testify/assert"
	"golang_boilerplate/testutil"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"errors"
	"golang_boilerplate/instrumentation"
	"net/http/httptest"
)

func TestCreatePostInDB(t *testing.T) {
	db, mock := testutil.SetupMockTestDB(t)
	defer testutil.CloseTestDB(db)
	logger := testutil.TestLogger()
	w := httptest.NewRecorder()
	txn := instrumentation.GetNewRelicTransaction(w, false)

	row := sqlmock.NewRows([]string{"id"}).AddRow("ID1")
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT into posts`).WithArgs("Why am I awesome?", "Just because Im awesome").WillReturnRows(row)
	mock.ExpectCommit()

	postsRepository := NewPostsRepository(db, logger)

	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	id, err := postsRepository.Create(post, txn)

	assert.NoError(t, err, "Successfully Created a Post")
	assert.NotEmpty(t, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestErrorWhileBeginningTxn(t *testing.T) {
	db, mock := testutil.SetupMockTestDB(t)
	defer testutil.CloseTestDB(db)
	logger := testutil.TestLogger()
	w := httptest.NewRecorder()
	txn := instrumentation.GetNewRelicTransaction(w, false)

	mock.ExpectBegin().WillReturnError(errors.New("Begin Txn Failed"))

	postsRepository := NewPostsRepository(db, logger)

	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	id, err := postsRepository.Create(post, txn)

	assert.Error(t, err, "Error while committing to DB")
	assert.Empty(t, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestErrorWhileCreatingPostInDB(t *testing.T) {
	db, mock := testutil.SetupMockTestDB(t)
	defer testutil.CloseTestDB(db)
	logger := testutil.TestLogger()
	w := httptest.NewRecorder()
	txn := instrumentation.GetNewRelicTransaction(w, false)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT into posts`).WithArgs("Why am I awesome?", "Just because Im awesome").WillReturnError(errors.New("DB Error"))
	mock.ExpectRollback()

	postsRepository := NewPostsRepository(db, logger)

	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	id, err := postsRepository.Create(post, txn)

	assert.Error(t, err, "Error while communicating to DB")
	assert.Empty(t, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestErrorWhileCommittingPostInDB(t *testing.T) {
	db, mock := testutil.SetupMockTestDB(t)
	defer testutil.CloseTestDB(db)
	logger := testutil.TestLogger()
	w := httptest.NewRecorder()
	txn := instrumentation.GetNewRelicTransaction(w, false)

	row := sqlmock.NewRows([]string{"id"}).AddRow("ID1")
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT into posts`).WithArgs("Why am I awesome?", "Just because Im awesome").WillReturnRows(row)
	mock.ExpectCommit().WillReturnError(errors.New("Commit Failed"))

	postsRepository := NewPostsRepository(db, logger)

	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	id, err := postsRepository.Create(post, txn)

	assert.Error(t, err, "Error while committing to DB")
	assert.Empty(t, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

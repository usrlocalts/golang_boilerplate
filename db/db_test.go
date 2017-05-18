package db

import (
	"testing"
	"golang_boilerplate/testutil"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestNewDB(t *testing.T) {
	logger := testutil.TestLogger()
	db := NewDB(logger, testutil.TestDBUrl, 5)
	defer db.Close()
	assert.Equal(t, reflect.TypeOf(db).String(), "*sqlx.DB")
}

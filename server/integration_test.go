package server


import (
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
	"testing"
	"golang_boilerplate/testutil"
	"net/http"
	"golang_boilerplate/posts_creation"
	"github.com/stretchr/testify/assert"
	"golang_boilerplate/appcontext"
	"golang_boilerplate/config"
	"golang_boilerplate/errors"
)

func TestSuccessfulPostsCreation(t *testing.T) {

	logger := testutil.TestLogger()
	config := config.Load()
	ctx := appcontext.NewAppContext(logger, config)
	db := testutil.SetupTestDB(logger)
	router := Router(ctx, db)
	ts := testutil.NewTestServer(router)
	defer ts.Close()
	defer db.Close()

	requestBody, _ := json.Marshal(&posts_creation.PostsCreationRequest{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	})

	hc := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/posts", ts.URL), bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")

	testutil.WithCleanTable(func() {
		res, _ := hc.Do(req)
		data, _ := ioutil.ReadAll(res.Body)
		reply := posts_creation.PostsCreationResponse{}
		err := json.Unmarshal(data, &reply)
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusCreated)
		assert.Equal(t, reply.Post.Topic, "Why am I awesome?")
		assert.Equal(t, reply.Post.Body, "Just because Im awesome")
		assert.NotEmpty(t, reply.Post.Id)
	}, db, "posts")

}

func TestPostsCreationError(t *testing.T) {

	logger := testutil.TestLogger()
	config := config.Load()
	ctx := appcontext.NewAppContext(logger, config)
	db := testutil.SetupTestDB(logger)
	router := Router(ctx, db)
	ts := testutil.NewTestServer(router)
	defer ts.Close()
	db.Close()

	requestBody, _ := json.Marshal(&posts_creation.PostsCreationRequest{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	})

	hc := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/posts", ts.URL), bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")

	testutil.WithCleanTable(func() {

		res, err := hc.Do(req)
		assert.NoError(t, err)
		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		reply := posts_creation.PostsCreationResponse{}
		err = json.Unmarshal(data, &reply)
		assert.NoError(t, err)

		fmt.Println(reply)

		assert.Len(t, reply.Errors, 1)
		assert.Equal(t, reply.Errors[0].Code, errors.GenericServiceError)
		assert.Equal(t, res.StatusCode, http.StatusInternalServerError)
		assert.Empty(t, reply.Post)
	}, db, "posts")

}

func TestPostsCreationMalformedError(t *testing.T) {

	logger := testutil.TestLogger()
	config := config.Load()
	ctx := appcontext.NewAppContext(logger, config)
	db := testutil.SetupTestDB(logger)
	router := Router(ctx, db)
	ts := testutil.NewTestServer(router)
	defer ts.Close()
	defer db.Close()


	requestBody := []byte("aksjdhf")

	hc := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/posts", ts.URL), bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")

	testutil.WithCleanTable(func() {

		res, err := hc.Do(req)
		assert.NoError(t, err)
		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		reply := posts_creation.PostsCreationResponse{}
		err = json.Unmarshal(data, &reply)
		assert.NoError(t, err)

		assert.Equal(t, res.StatusCode, http.StatusBadRequest)
		assert.Len(t, reply.Errors, 1)
		assert.Equal(t, reply.Errors[0].Code, errors.MalformedJSON)
		assert.Empty(t, reply.Post)
	}, db, "posts")

}

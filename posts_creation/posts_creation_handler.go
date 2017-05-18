package posts_creation

import (
	"encoding/json"
	e "golang_boilerplate/errors"
	"golang_boilerplate/handler"
	"golang_boilerplate/logger"
	"net/http"
	"golang_boilerplate/instrumentation"
	"github.com/newrelic/go-agent"
	"github.com/pkg/errors"
)

func PostsCreationHandler(postsService PostsService, logger logger.Log, newrelicEnabled bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postsCreationRequest PostsCreationRequest

		postsCreationResponse := &PostsCreationResponse{}
		err := json.NewDecoder(r.Body).Decode(&postsCreationRequest)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			logger.Warn(err.Error())
			resp := NewErrorPostCreationResponse([]*e.Error{e.NewError(e.MalformedJSON, errors.New("Malformed JSON"))})
			handler.RespondWithCustomErrors(w, resp, http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		txn := instrumentation.GetNewRelicTransaction(w, newrelicEnabled)
		segment := newrelic.StartSegment(txn, "createPosts")

		post := NewPost(postsCreationRequest.Topic, postsCreationRequest.Body)
		createdPostId, err := postsService.Create(post, txn)

		segment.End()

		if err != nil {
			logger.Errorf("error creating posts; %v", err)
			resp := NewErrorPostCreationResponse([]*e.Error{e.NewError(e.GenericServiceError, errors.New("Failed to create Post"))})
			handler.RespondWithCustomErrors(w, resp, http.StatusInternalServerError)
			return
		}

		postsCreationResponse = NewPostsCreationResponse(post, createdPostId)

		w.WriteHeader(http.StatusCreated)

		response, _ := json.Marshal(postsCreationResponse)
		w.Write(response)
		logger.Info("Successfully created post with id : ", createdPostId)
	}
}

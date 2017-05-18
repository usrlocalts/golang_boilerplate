package posts_creation

import (
	"golang_boilerplate/logger"
	"github.com/newrelic/go-agent"
)

type PostsService interface {
	Create(*Post, newrelic.Transaction) (string, error)
}

type PostsServiceImpl struct {
	postsRepository PostsRepository
	logger          logger.Log
}

func NewPostsService(postsRepository PostsRepository, logger logger.Log) PostsService {
	return &PostsServiceImpl{
		postsRepository: postsRepository,
		logger:          logger,
	}
}

func (t *PostsServiceImpl) Create(post *Post, txn newrelic.Transaction) (string, error) {
	id, err := t.postsRepository.Create(post, txn)
	if err != nil {
		t.logger.Errorf("Failed to create post: %v", err)
		return "", err
	}

	return id, nil
}

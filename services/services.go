package services

import (
	"golang_boilerplate/posts_creation"
	"github.com/jmoiron/sqlx"
	"golang_boilerplate/logger"
)

type Services struct {
	PostsService posts_creation.PostsService
}

func New(logger logger.Log, db *sqlx.DB) *Services {
	postsRepository := posts_creation.NewPostsRepository(db, logger)
	postsService := posts_creation.NewPostsService(postsRepository, logger)
	return &Services{PostsService: postsService}
}

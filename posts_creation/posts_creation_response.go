package posts_creation

import "golang_boilerplate/errors"

type PostsDescription struct {
	Id    string `json:"id"`
	Body string `json:"body"`
	Topic  string `json:"topic"`
}

type PostsCreationResponse struct {
	Post   *PostsDescription `json:"post"`
	Errors []*errors.Error   `json:"errors"`
}

func NewPostsCreationResponse(post *Post, id string) *PostsCreationResponse {
	postsDescription := &PostsDescription{Id: id, Topic: post.Topic, Body: post.Body}
	return &PostsCreationResponse{Post: postsDescription}
}

func NewErrorPostCreationResponse(errors []*errors.Error) *PostsCreationResponse {
	return &PostsCreationResponse{
		Errors: errors,
	}
}

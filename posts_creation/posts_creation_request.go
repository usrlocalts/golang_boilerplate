package posts_creation

type PostsCreationRequest struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
}
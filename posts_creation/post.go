package posts_creation

type Post struct {
	Topic string
	Body  string
}

func NewPost(topic, body string) *Post {
	return &Post{
		Topic: topic,
		Body:  body,
	}
}

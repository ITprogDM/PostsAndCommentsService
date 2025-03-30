package models

import "time"

type InputComment struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Post    int    `json:"post"`
	ReplyTo *int   `json:"replyTo,omitempty"`
}

type InputPost struct {
	Name            string `json:"name"`
	Content         string `json:"content"`
	Author          string `json:"author"`
	CommentsAllowed bool   `json:"commentsAllowed"`
}

type Mutation struct {
}

type PostGraph struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
}

type Query struct {
}

type Subscription struct {
}

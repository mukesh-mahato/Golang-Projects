package main

import "math/rand"

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Description string `json:"description"`
}

func NewBook(title, author, published, description string) *Book {
	return &Book{
		ID:          rand.Intn(10000),
		Title:       title,
		Author:      author,
		Published:   published,
		Description: description,
	}
}

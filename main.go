package main

import (
	"time"
)

type Book struct {
	id string `json:"id"`
	title string `json:"title"`
	description string `json:"description"`
	author string 	`json:"author"`
	genre string 	`json:"genre"`
	publication_date string `json:"publication_date"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

func main() {
	println("Hello, World!")
}
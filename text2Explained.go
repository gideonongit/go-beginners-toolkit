package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("📝 Simple API Client Example")
	fmt.Println("============================")
	
	posts, err := fetchPosts()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Display first 3 posts
	for i := 0; i < 3 && i < len(posts); i++ {
		fmt.Printf("\nPost %d:\n", i+1)
		fmt.Printf("Title: %s\n", posts[i].Title)
		fmt.Printf("Body: %s\n", posts[i].Body[:50]+"...") // First 50 chars
		fmt.Println("---")
	}
}

func fetchPosts() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}
	
	return posts, nil
}
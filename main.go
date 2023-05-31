package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Post struct {
	Message string `json:"message"`
}

func main() {
	accessToken := os.Getenv("FACEBOOK_ACCESS_TOKEN")
	pageID := os.Getenv("FACEBOOK_PAGE_ID")
	url := fmt.Sprintf("https://graph.facebook.com/v13.0/%s/feed", pageID)

	post := Post{
		Message: os.Getenv("FACEBOOK_POST_MESSAGE")),
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.URL.RawQuery = fmt.Sprintf("access_token=%s", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Post successfully published on Facebook!")
}

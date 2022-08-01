package main

import (
	"fmt"
	"net/http"
)

func main() {
	contentType, error := contentType("https://linkedin.com")
	if error != nil {
		fmt.Printf("ERROR: %s\n", error)
	} else {
		fmt.Println(contentType)
	}
}

//Return the value of Content-Type header returned by
// making an HTTP request to url
func contentType(url string) (string, error) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Printf("ERROR: %s\n", error)
		return "", error
	}

	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("Content-Type not present")
	}
	return contentType, error
}

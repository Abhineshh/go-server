package main

import (
	"fmt"
	"io"
	//"log"
	"net/http"
	"os"
)

func main() {
    url := "https://www.gatesnotes.com/"

    // Send a GET request to the specified URL
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // Create a new file to store the response
    file, err := os.Create("ging.html")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Copy the response body to the file
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println("HTML content has been written to response.html")
}

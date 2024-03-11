package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	c := openai.NewClient("openai api key") //the openai api key

	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Prompt > ")
	if !s.Scan() {
		panic("failed to get user input")
	}
	req := openai.ImageRequest{
		Prompt:         s.Text(),
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}
	resp, err := c.CreateImage(context.Background(), req)

	if err != nil {
		panic(err)
	}
	b, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		panic(err)
	}

}

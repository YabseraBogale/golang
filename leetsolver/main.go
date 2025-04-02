package main

import (
	"context"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/otiai10/gosseract/v2"
	"github.com/vova616/screenshot"
	"google.golang.org/api/option"
)

func main() {
	api := os.Getenv("api")
	ctx := context.Background()
	tesract := gosseract.NewClient()
	defer tesract.Close()
	client, err := genai.NewClient(ctx, option.WithAPIKey(api))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	file, err := os.Create("image.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	screen, err := screenshot.CaptureScreen()
	if err != nil {
		log.Println(err)
	}
	err = png.Encode(file, screen)
	if err != nil {
		log.Println(err)
		return
	}
	err = tesract.SetImage(file.Name())
	if err != nil {
		log.Println(err)
	}
	text, err := tesract.Text()
	if err != nil {
		log.Println(err)
	}

	question := ""
	if question == "" {
		fmt.Println(text)
	} else {
		model := client.GenerativeModel("gemini-2.0-flash")
		resp, err := model.GenerateContent(ctx, genai.Text("How does AI work?"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
	}

}

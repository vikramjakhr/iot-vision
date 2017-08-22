package services

import (
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"

	"os"
	"fmt"
)

type TextReco struct {
	Url          string
	Text         string
	FaceInfoList []FaceInfo
}

var TextRecoChan chan TextReco = make(chan TextReco, 10)

func Process(file, object string) {
	var text string = "No text found"
	face := make(chan []FaceInfo)
	url := make(chan string)
	go func() {
		SaveToCloudStorage(file, object, url)
	}()
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		SearchFaces("mycollection", image.Content, face)
	}()

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		fmt.Println(err)
	}

	if len(annotations) != 0 {
		text = ""
		for _, annotation := range annotations {
			text += annotation.Description + "<br />"
		}
	}
	tr := TextReco{
		Url:          <-url,
		Text:         text,
		FaceInfoList: <-face,
	}
	close(face)
	close(url)
	TextRecoChan <- tr
	client.Close()
}

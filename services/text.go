package services

import (
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"

	"os"
	"fmt"
)

type TextReco struct {
	Url         string
	Text        string
	MatchedUrls []string
}

var TextRecoChan chan TextReco = make(chan TextReco, 10)

func Process(file, object string) {
	var url string
	var text string = "No text found"
	url = SaveToCloudStorage(file, object)
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

	MatchedUrls := SearchFaces("mycollection", image.Content)

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		fmt.Println(err)
	}

	tr := TextReco{
		Url:         url,
		Text:        text,
		MatchedUrls: MatchedUrls,
	}
	if len(annotations) != 0 {
		for _, annotation := range annotations {
			text += annotation.Description + "<br />"
		}
		tr.Text = text
	}
	TextRecoChan <- tr
	client.Close()
}

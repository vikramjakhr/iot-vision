package services

import (
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"

	"os"
	"fmt"
)

type TextReco struct {
	Url  string
	Text string
}

var TextRecoChan chan TextReco = make(chan TextReco, 10)

func DetectText(file, object string) {
	var url string
	var text string = ""
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
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		fmt.Println(err)
	}

	if len(annotations) == 0 {
		fmt.Println("No text found.")
	} else {
		for _, annotation := range annotations {
			text += annotation.Description + "<br />"
		}
		tr := TextReco{
			Url:  url,
			Text: text,
		}
		fmt.Println("Text:", tr)
		TextRecoChan <- tr
	}
	client.Close()
}

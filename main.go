package main

import (
	_ "gitlab.intelligrape.net/tothenew/vision/routers"
	"github.com/astaxie/beego"

	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"

	"os"
	"fmt"
)

func main() {
	beego.Run()
	detectText("/home/vikram/Desktop/img/1.jpg")
}

func detectText(file string) {
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
		fmt.Println("Text:")
		for _, annotation := range annotations {
			fmt.Printf(annotation.Description)
		}
	}
}

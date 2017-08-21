package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"errors"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"fmt"
)

func getSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewEnvCredentials(),
	})
	log.Println("#### Session created")
	if err != nil {
		errors.New("Error while creating session")
	}
	return sess
}

func CreateCollection(name string) *rekognition.CreateCollectionOutput {
	client := rekognition.New(getSession())
	input := &rekognition.CreateCollectionInput{
		CollectionId: &name,
	}
	req, resp := client.CreateCollectionRequest(input)
	err := req.Send()
	if err == nil {
		fmt.Println(resp)
	}
	return resp
}

func IndexFaces(collName, imgName string) *rekognition.IndexFacesOutput {
	client := rekognition.New(getSession())
	bucket := "ttn-aws-iot"
	s3Object := &rekognition.S3Object{
		Name:   &imgName,
		Bucket: &bucket,
	}
	image := &rekognition.Image{
		S3Object: s3Object,
	}
	input := &rekognition.IndexFacesInput{
		CollectionId: &collName,
		Image:        image,

	}
	req, resp := client.IndexFacesRequest(input)
	err := req.Send()
	if err == nil {
		fmt.Println(resp)
	}
	return resp
}

func SearchFaces(collName string, bts []byte) []string {
	var s []string = []string{}
	client := rekognition.New(getSession())
	image := &rekognition.Image{
		Bytes: bts,
	}
	fmt.Println(len(bts))
	input := &rekognition.SearchFacesByImageInput{
		CollectionId: &collName,
		Image:        image,
	}
	req, resp := client.SearchFacesByImageRequest(input)
	err := req.Send()
	if err == nil {
		fmt.Println(resp)
		for _, face := range resp.FaceMatches {
			s = append(s, "https://s3.amazonaws.com/ttn-aws-iot/" + *face.Face.FaceId+".jpg")
		}
	} else {
		fmt.Println(err)
	}
	return s
}

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

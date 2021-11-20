package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	// Core Golang Dependencies.
	"context"
	"fmt"
	"log"
)

func main()  {

	PutItem()
	GetItem()
}


func PutItem(){
	// Contex
	ctx := context.Background()

	// Table Name
	tableName := "MyTable"

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	type Item struct {
		Id   string`json:"id"`
		Name string`json:"name"`
	}
	// Item to send to the DB.
	item := Item {
		Id: "1234567890000",
		Name: "Alexandro4",
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("unable to marsh the item")
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(ctx, input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Println("Successfully added '" + item.Name + " to table '" + tableName)
}

func GetItem(){
	// Contex
	ctx := context.Background()

	// Table Name
	tableName := "MyTable"
	id := "1234567890000"
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	type Item struct {
		Id   string`json:"id"`
		Name string`json:"name"`
	}


	output := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
			},
			TableName: aws.String(tableName),
		}


	_, err = svc.GetItem(ctx, output)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	item := Item{}
		err = attributevalue.UnmarshalMap(output.Key, &item)
		if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
		}

	fmt.Println(" Name retrieved: " + item.Name)
	}






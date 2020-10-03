package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://daityaADMs:9Nmfbp6l1MAz@cluster0.7tvor.mongodb.net/<dbname>?retryWrites=true&w=majority",
	))
	fmt.Println()
	if err != nil { log.Fatal(err) }


}

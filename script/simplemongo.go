package script

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Subject string
	Title   string
	Watched int
}

func ConnectMongo() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	users := client.Database("playground").Collection("users")
	// res, err := users.InsertOne(context.Background(), Person{Subject: "logic", Watched: 200, Title: "learn algorithm"})
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	// fmt.Println(res.InsertedID)
	// res, err := users.Find(context.Background(), bson.D{{"subject", "logic"}})
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	// var results []bson.M
	//
	//	if err = res.All(context.Background(), &results); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	for _, result := range results {
	//		fmt.Println(result)
	//	}

	ctx := context.Background()
	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
    [
		{
			"$match": {"title":{ "$regex": "how", "$options": "i" }}
		},
		{
			"$group": {
					"_id": "$subject",
					"totalWatched": { "$sum": "$watched" },
					"count": { "$sum": 1 }
			}
		}
    ]
	`)), true, &pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	csr, err := users.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	fmt.Println(result)
	// Disconnect from MongoDB
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB.")
}

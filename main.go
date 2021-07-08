package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	api "github.com/suhani-kohli/calorie-tracker-service/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	router := mux.NewRouter()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("calorie-tracker").Collection("diary")

	fmt.Printf(" %T", collection)
	fmt.Printf(" %T", ctx)

	router.HandleFunc("/foodlogs", api.CreateFoodLog(collection)).Methods(http.MethodPost)
	router.HandleFunc("/foodlogs", api.GetFoodLogs(collection)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))

}

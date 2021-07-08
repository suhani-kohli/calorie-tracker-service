package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	router := mux.NewRouter()
	// host := "localhost"
	// port := 5432
	// user := "suhani1"
	// password := "password"
	// dbname := "gotest"

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// fmt.Println(psqlInfo)
	// db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

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

	// router.HandleFunc("/item", api.GetItem(client)).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))

}

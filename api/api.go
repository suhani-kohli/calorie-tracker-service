package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodLog struct {
	MealName string `json:"mealName,omitempty" bson:"mealName,omitempty"`
	FoodItem string `json:"foodItem,omitempty" bson:"foodItem,omitempty"`
	Date     string `json:"date,omitempty" bson:"date,omitempty"`
}

func CreateFoodLog(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println("body", reqBody)
		var foodLog FoodLog
		err := json.Unmarshal(reqBody, &foodLog)
		if err != nil {
			fmt.Printf("There was an error decoding the json. err = %s", err)
			return
		}
		fmt.Println("food diary", foodLog)

		result, err := collection.InsertOne(context.TODO(), foodLog)

		if err != nil {

			return
		}
		fmt.Println(result)

		json.NewEncoder(w).Encode(result)
	}
}

func GetFoodLogs(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query()["date"][0]
		var foodLogs []FoodLog
		filter := bson.D{{"date", date}}

		cur, err := collection.Find(context.TODO(), filter)
		if err = cur.All(context.Background(), &foodLogs); err != nil {
			//   log.Fatal(err)
			fmt.Printf("There was an error while parsing result. err = %s", err)
			return

		}
		fmt.Println(foodLogs)
		json.NewEncoder(w).Encode(foodLogs)
	}

}

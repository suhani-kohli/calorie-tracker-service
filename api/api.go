package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type FoodDiary struct {
	MealName string `json:"mealName,omitempty" bson:"mealName,omitempty"`
	FoodItem string `json:"foodItem,omitempty" bson:"foodItem,omitempty"`
	Date     string `json:"date,omitempty" bson:"date,omitempty"`
}

func CreateItem(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println("body", reqBody)
		var foodDiary FoodDiary
		err := json.Unmarshal(reqBody, &foodDiary)
		if err != nil {
			fmt.Printf("There was an error decoding the json. err = %s", err)
			return
		}
		fmt.Println("food diary", foodDiary)

		result, err := collection.InsertOne(context.TODO(), foodDiary)

		if err != nil {

			return
		}
		fmt.Println(result)

		json.NewEncoder(w).Encode(result)
	}
}

package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type FoodDiary struct {
	MealName string `json:"_mealName,omitempty" bson:"_mealName,omitempty"`
	FoodItem string `json:"_foodItem,omitempty" bson:"_foodItem,omitempty"`
	Date     string `json:"_date,omitempty" bson:"_date,omitempty"`
}

func CreateItem(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var foodDiary FoodDiary
		json.Unmarshal(reqBody, &foodDiary)

		result, err := collection.InsertOne(context.TODO(), foodDiary)

		if err != nil {

			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

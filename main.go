package main

import (
	"fmt"

	"recommendation-system-go/engine"
	"recommendation-system-go/model"
)

func main() {

	ratings := model.UserRatings{

		"UserA": {
			"Pod": 1,
			"Job": 1,
		},

		"UserB": {
			"Pod":       1,
			"Job":       1,
			"ConfigMap": 1,
		},

		"UserC": {
			"Pod":    1,
			"Secret": 1,
		},
	}

	recommender := engine.NewRecommender(ratings)

	recommendations := recommender.Recommend("UserA")

	fmt.Println("Recommendations for UserA:")

	for _, rec := range recommendations {

		fmt.Printf(
			"Item: %s (similarity score %.2f)\n",
			rec.Item,
			rec.Score,
		)
	}
}
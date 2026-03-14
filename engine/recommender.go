package engine

import (
	"sort"

	"recommendation-system-go/model"
	"recommendation-system-go/similarity"
)

// Recommendation represents a recommended item and its score.
type Recommendation struct {
	Item  model.ItemID
	Score float64
}

// Recommender contains the dataset used to compute recommendations.
type Recommender struct {
	Ratings model.UserRatings
}

// NewRecommender creates a new recommendation engine.
func NewRecommender(ratings model.UserRatings) *Recommender {
	return &Recommender{
		Ratings: ratings,
	}
}

// Recommend returns items that a user may like based on similar users.
func (r *Recommender) Recommend(user model.UserID) []Recommendation {

	targetRatings := r.Ratings[user]

	var bestUser model.UserID
	bestSimilarity := 0.0

	for otherUser, otherRatings := range r.Ratings {

		if otherUser == user {
			continue
		}

		sim := similarity.CosineSimilarity(targetRatings, otherRatings)

		if sim > bestSimilarity {
			bestSimilarity = sim
			bestUser = otherUser
		}
	}

	if bestUser == "" {
		return nil
	}

	var recommendations []Recommendation

	for item, rating := range r.Ratings[bestUser] {

		if targetRatings[item] == 0 && rating > 0 {

			recommendations = append(recommendations, Recommendation{
				Item:  item,
				Score: bestSimilarity,
			})
		}
	}

	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	return recommendations
}
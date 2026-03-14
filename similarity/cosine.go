package similarity

import (
	"math"

	"recommendation-system-go/model"
)

// CosineSimilarity calculates how similar two users are based on their ratings.
//
// Formula:
//
// similarity = dot(A,B) / (|A| * |B|)
//
// A = user1 rating vector
// B = user2 rating vector
//
// The result ranges between 0 and 1
// 1 = identical preferences
// 0 = completely different
func CosineSimilarity(
	a map[model.ItemID]model.RatingValue,
	b map[model.ItemID]model.RatingValue,
) float64 {

	var dotProduct float64
	var magnitudeA float64
	var magnitudeB float64

	for item, ratingA := range a {

		ratingB := b[item]

		dotProduct += float64(ratingA * ratingB)

		magnitudeA += float64(ratingA * ratingA)
		magnitudeB += float64(ratingB * ratingB)
	}

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB))
}
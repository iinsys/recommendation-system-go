package model

// UserID represents a user in the system.
type UserID string

// ItemID represents an item that can be recommended.
type ItemID string

// RatingValue represents how strongly a user interacted with an item.
// In this simple example we use 1 for "interacted" and 0 for "no interaction".
type RatingValue float64

// UserRatings represents a matrix:
//
// user -> item -> rating
//
// Example:
// UserA -> {Pod:1, Job:1}
// UserB -> {Pod:1, Job:1, ConfigMap:1}
type UserRatings map[UserID]map[ItemID]RatingValue
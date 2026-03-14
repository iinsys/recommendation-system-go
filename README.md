# recommendation-system-go

A collaborative filtering recommendation engine written in Go. It uses cosine similarity to find users with similar preferences and recommends items they have interacted with that the target user has not yet seen.

## How it works

1. Each user has a rating vector — a map of items to rating values (e.g. 1 for interacted, 0 for not).
2. The engine computes cosine similarity between the target user and all other users.
3. The most similar user is identified, and items they rated that the target user has not seen are returned as recommendations, scored by the similarity value.

## Getting started

**Prerequisites:** Go 1.18 or later

```bash
git clone https://github.com/iinsys/recommendation-system-go.git
cd recommendation-system-go
go mod tidy
go run main.go
```

## Example

```go
ratings := model.UserRatings{
    "UserA": {"Pod": 1, "Job": 1},
    "UserB": {"Pod": 1, "Job": 1, "ConfigMap": 1},
    "UserC": {"Pod": 1, "Secret": 1},
}

recommender := engine.NewRecommender(ratings)
recommendations := recommender.Recommend("UserA")
```

Output:

```
Recommendations for UserA:
Item: ConfigMap (similarity score 1.00)
```

## Cosine similarity

The similarity between two users is computed as:

```
similarity = dot(A, B) / (|A| * |B|)
```

Where A and B are the rating vectors of two users. The result is a value between 0 and 1, where 1 means identical preferences and 0 means no overlap.

## License

See [LICENSE](./LICENSE).

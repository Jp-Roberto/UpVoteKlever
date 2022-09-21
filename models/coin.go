package models

import "github.com/google/uuid"

type Vote struct {
	Crypt string `json:"crypt" bson:"crypt`
	Vote  int    `json:"vote" bson:"vote"`
}

type VoteBody struct {
	Vote int `json:"vote" bson:"vote"`
}

type CoinBody struct {
	Name  string `json:"name" bson:"name"`
	Code  string `json:"code" bson:"code"`
	Votes []Vote `json:"votes" bson:"votes"`
}

type CoinVotes struct {
	Name  string `json:"name" bson:"name"`
	Code  string `json:"code" bson:"code"`
	Votes int    `json:"votes" bson:"votes"`
}

type Coin struct {
	Id    uuid.UUID `json:"id,omitempty" bson:"_id"`
	Name  string    `json:"name" bson:"name"`
	Code  string    `json:"code" bson:"code"`
	Votes []Vote    `json:"votes" bson:"votes"`
}

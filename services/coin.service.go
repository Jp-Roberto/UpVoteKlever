package services

import (
	"github.com/Jp-Roberto/challengerklever/models"
	uuid "github.com/satori/go.uuid"
)

type CoinService interface {
	CreateCoin(*models.Coin) error
	GetAllVotes() ([]*models.CoinVotes, error)
	HandleVote(coinId uuid.UUID, userId string, vote int) (*int, error)
}

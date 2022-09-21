package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Jp-Roberto/challengerklever/models"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoinServiceImpl struct {
	coincollection *mongo.Collection

	ctx context.Context
}

func NewCoinService(coincollection *mongo.Collection, ctx context.Context) CoinService {
	return &CoinServiceImpl{
		coincollection: coincollection,
		ctx:            ctx,
	}
}

func (c *CoinServiceImpl) resetVotes() ([]models.Vote, error) {
	var votes []models.Vote

	cursor, err := c.coincollection.Find(c.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next((c.ctx)) {
		var crypt models.Coin
		var vote models.Vote
		err := cursor.Decode(&crypt)
		if err != nil {
			return nil, err
		}
		vote.Crypt = "Crypto" // lembrar de alterar no termino para a chamada da func correta!
		vote.Vote = 0
		votes = append(votes, vote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(votes) == 0 {
		return nil, errors.New(("Crypto not found"))
	}

	return votes, err
}

func (c *CoinServiceImpl) CreateCoin(coin *models.Coin) error {

	votes, err := c.resetVotes()
	if err != nil {
		return err
	}
	coin.Votes = votes
	_, err = c.coincollection.InsertOne(c.ctx, coin)
	return err

}

func (c *CoinServiceImpl) GetAllVotes() ([]*models.CoinVotes, error) {
	var coins []*models.CoinVotes

	cursor, err := c.coincollection.Find(c.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next((c.ctx)) {
		var crypt models.Coin
		var coinVotes models.CoinVotes
		err := cursor.Decode(&crypt)
		if err != nil {
			return nil, err
		}
		fmt.Println("CryptId: ", crypt.Id)
		//ATENÇÃO!!!!
		// numberOfVotes, err := c.getNumberOfVotes(crypt.Id)  // Lembrar de tirar comentário e fazer o User.ID com uuuid  e alterar os campos crypt
		// if err != nil {
		// 	fmt.Println("error: ", err)
		// 	return nil, err
		// }
		// coinVotes.Votes = *numberOfVotes
		coins = append(coins, &coinVotes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(coins) == 0 {
		return nil, errors.New(("NOTHING HAS BEEN FOUND!!!"))
	}

	return coins, nil
}

func (c *CoinServiceImpl) getNumberOfVotes(coinId uuid.UUID) (*int, error) {
	var votes = 0
	var coin models.Coin

	err := c.coincollection.FindOne(c.ctx, bson.D{bson.E{Key: "_id", Value: coinId}}).Decode(coin)
	if err != nil {
		return nil, err
	}

	for _, s := range coin.Votes {
		votes += s.Vote
	}
	fmt.Println("TOTAL VOTES: ", votes)
	return &votes, err
}

func (c *CoinServiceImpl) HandleVote(coinId uuid.UUID, userId string, vote int) (*int, error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: coinId},
		bson.E{Key: "votes", Value: bson.D{
			bson.E{
				Key: "$elemMatch", Value: bson.E{
					Key: "user_id", Value: userId,
				},
			},
		}},
	}
	query := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "votes.$.vote", Value: 1},
		}},
	}
	result, _ := c.coincollection.UpdateOne(c.ctx, filter, query)

	if result.MatchedCount != 1 {
		return nil, errors.New("NO CRYPTOMED FOUND FOR UPDATE")
	}

	numberOfVotes, err := c.getNumberOfVotes(coinId)

	if err != nil {
		return nil, err
	}

	return numberOfVotes, nil
}

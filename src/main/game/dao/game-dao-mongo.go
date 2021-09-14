package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/game/model"
)

type GameDaoMongo struct{
	C *mongo.Collection
}

func (dao GameDaoMongo) CreateGame(game model.Game) (model.Game, error) {
	_, err := dao.C.InsertOne(context.TODO(), game)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted a single document: %+v\n", game)
	return game, nil
}

func (dao GameDaoMongo) UpdateGame(game model.Game) (model.Game, error) {
	filter := bson.D{{"_id", game.ID}}

	updateResult, err := dao.C.UpdateOne(context.TODO(), filter, game)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return game, err
}

func (dao GameDaoMongo) GetGame(gameid string) (model.Game, error) {
	filter := bson.D{{"_id", gameid}}

	var result model.Game

	err := dao.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, err
}



func (dao GameDaoMongo) GetGamesByWeekOrderByGameStart(weekid string) ([]model.Game, error) {
	panic("implement me")
}

func (dao GameDaoMongo) GetGamesByWeek(weekid string) ([]model.Game, error) {
	findOptions := options.Find()
	//filter := bson.D{{"username", seasonid}}

	var results []model.Game
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Game
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		if elem.Weekid == weekid {
			results = append(results, elem)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	err = cur.Close(context.TODO())
	if err != nil {
		panic(err)
	}

	return results,err
}

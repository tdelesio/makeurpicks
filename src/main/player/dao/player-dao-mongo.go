package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/player/model"
)

type PlayerDaoMongo struct {
	C *mongo.Collection
}

func (dao PlayerDaoMongo) CreatePlayer(player model.Player) (model.Player, error) {
	_, err := dao.C.InsertOne(context.TODO(), player)
	if err != nil {
		panic(err)
	}

	//player.Username = insertedResult.InsertedID.(primitive.ObjectID)

	fmt.Printf("Inserted a single document: %+v\n", player)
	return player, nil
}

func (dao PlayerDaoMongo) UpdatePlayer(player model.Player) (model.Player, error) {
	filter := bson.D{{"_id", player.Username}}

	updateResult, err := dao.C.UpdateOne(context.TODO(), filter, player)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return player, err
}

func (dao PlayerDaoMongo) GetPlayer(username string) (model.Player, error) {
	filter := bson.D{{"_id", username}}

	var result model.Player

	err := dao.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, err
}

func (dao PlayerDaoMongo) GetPlayers() (*[]model.Player, error) {
	findOptions := options.Find()

	var results []model.Player
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Player
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}


		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	err = cur.Close(context.TODO())
	if err != nil {
		panic(err)
	}

	return &results,err
}




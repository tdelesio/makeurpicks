package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/season/model"
)

type SeasonDaoMongo struct {
	C *mongo.Collection
}

func (dao SeasonDaoMongo) GetSeasonsByLeagueType(leagyetype string) (*[]model.Season, error) {
	findOptions := options.Find()

	var results []model.Season
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Season
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		if elem.LeagueType == leagyetype {
			results = append(results, elem)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return &results,err
}

func (dao SeasonDaoMongo) CreateSeason(season model.Season) (*model.Season, error) {
	insertedResult, err := dao.C.InsertOne(context.TODO(), season)
	if err != nil {
		return nil, err
	}

	season.ID = insertedResult.InsertedID.(primitive.ObjectID)

	fmt.Println("Inserted a single document: ", insertedResult.InsertedID)

	return &season, nil
}

func (dao SeasonDaoMongo) UpdateSeason(season model.Season) (*model.Season,error) {
	filter := bson.D{{"_id", season.ID}}

	updateResult, err := dao.C.UpdateOne(context.TODO(), filter, season)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return &season, err
}

func (dao SeasonDaoMongo) DeleteSeason(id string) error {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	deleteResult, err := dao.C.DeleteOne(context.TODO(), bson.M{"_id": p})

	if err != nil {
		return err
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	return nil
}
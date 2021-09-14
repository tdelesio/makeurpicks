package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/league/model"
)

type LeaguerDaoMongo struct {
	C *mongo.Collection
}

func (dao LeaguerDaoMongo) CreateLeague(league model.League) (model.League, error) {
	_, err := dao.C.InsertOne(context.TODO(), league)
	if err != nil {
		panic(err)
	}

	//league.Username = insertedResult.InsertedID.(primitive.ObjectID)

	fmt.Printf("Inserted a single document: %+v\n", league)
	return league, nil
}

func (dao LeaguerDaoMongo) UpdateLeague(league model.League) (model.League, error) {
	filter := bson.D{{"_id", league.ID}}

	updateResult, err := dao.C.UpdateOne(context.TODO(), filter, league)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return league, err
}

func (dao LeaguerDaoMongo) GetLeague(id string) (model.League, error) {
	filter := bson.D{{"_id", id}}

	var result model.League

	err := dao.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, err
}

func (dao LeaguerDaoMongo) GetLeagueByLeagueName(leaguename string) (model.League, error) {
	filter := bson.D{{"leagueName", leaguename}}

	var result model.League

	err := dao.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, err
}


func (dao LeaguerDaoMongo) GetLeagues() ([]model.League,error) {
	findOptions := options.Find()

	var results []model.League
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.League
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

	return results,err
}

package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/team/model"
)

type TeamDaoMongo struct {
	//DB service.MongoDBConnI
	C *mongo.Collection
}

func (dao *TeamDaoMongo) GetTeam(id string) (model.Team, error) {
	filter := bson.D{{"_id", id}}

	var result model.Team

	err := dao.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, err
}

func (dao *TeamDaoMongo) CreateTeam(team model.Team) (*model.Team,error) {
	insertedResult, err := dao.C.InsertOne(context.TODO(), team)
	if err != nil {
		return nil, err
	}

	team.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &team, nil
}

func (dao *TeamDaoMongo) GetTeams(leagueType string) (*[]model.Team,error) {
	findOptions := options.Find()
	findOptions.SetLimit(32)

	var results []model.Team
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Team
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
	cur.Close(context.TODO())

	return &results,err
}
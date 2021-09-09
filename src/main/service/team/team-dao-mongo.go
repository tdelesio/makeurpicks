package team

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/model"
)

type TeamDaoMongo struct {
	//DB service.MongoDBConnI
	C *mongo.Collection
}

//func (dao *TeamDaoMongo) GetTeam(gameid string) (model.Team, error) {
//
//}

func (dao *TeamDaoMongo) CreateTeam(team model.Team) (*mongo.InsertOneResult,error) {
	return dao.C.InsertOne(context.TODO(), team)
}

func (dao *TeamDaoMongo) GetTeams(leagueType string) ([]*model.Team,error) {
	findOptions := options.Find()
	findOptions.SetLimit(32)

	var results []*model.Team
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Team
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results,err
}
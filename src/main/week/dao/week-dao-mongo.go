package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"makeurpicks/week/model"
)

type WeekDaoMongo struct{
	C *mongo.Collection
}

func (dao WeekDaoMongo) GetWeeksBySeason(seasonid string) ([]model.Week, error) {
	findOptions := options.Find()
	//filter := bson.D{{"username", seasonid}}

	var results []model.Week
	cur, err := dao.C.Find(context.TODO(),bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Week
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		if elem.SeasonId == seasonid {
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

func (dao WeekDaoMongo) CreateWeek(week model.Week) (model.Week, error) {
	_, err := dao.C.InsertOne(context.TODO(), week)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted a single document: %+v\n", week)
	return week, nil
}


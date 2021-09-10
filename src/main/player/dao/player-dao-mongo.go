package dao

import "go.mongodb.org/mongo-driver/mongo"

type PlayerDaoMongo struct {
	C *mongo.Collection
}

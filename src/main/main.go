package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"makeurpicks/pogs"
	team "makeurpicks/service/team"
	"net/http"
	"time"
)

const PICKEM = `pickem`
//const SUICIDE = `suicide`

const ProfileLocal = `LOCAL`
const ProfileDocker = `DOCKER`

const LocalDbHost = `localhost`
const LocalDbPort = 27017
const LocalDbName = `leagues`
const LocalSslMode = `disable`

const DockerDbHost = `apollo@aa-roachdb`
const DockerDbPort = 26257
const DockerDbName = `apollodb`
const DockerSslMode = `disable`

func main() {

	fmt.Println("Makeurpicks REST API")
	fmt.Println("initializing...")

	profile := ProfileLocal
	fmt.Printf("active profile: %s\n", profile)

	var DbData pogs.MongoMetadata

	switch profile {
	case ProfileLocal:
		DbData.Host = LocalDbHost
		DbData.Port = LocalDbPort
		DbData.SslMode = LocalSslMode
		DbData.DbName = LocalDbName
		break
	case ProfileDocker:
		DbData.Host = DockerDbHost
		DbData.Port = DockerDbPort
		DbData.SslMode = DockerSslMode
		DbData.DbName = DockerDbName
		break
	}


	fmt.Println("configuring the database (mongodb)...")
	uri := "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	//// Set client options
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//
	//// Connect to MongoDB
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Check the connection
	//err = client.Ping(context.TODO(), nil)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Connected to MongoDB!")

	fmt.Println("Configuring the DAOs...")
	teamDao := team.TeamDaoMongo{
		client.Database(DbData.DbName).Collection("teams"),
	}

	fmt.Println("Configuring the services...")
	teamService := team.TeamService{
		TeamRepository: &teamDao,
	}

	teamService.CreateAllTeams(PICKEM)

	fmt.Println("Configuring the static data...")
	teamMap := teamService.BuildTeamMap(PICKEM)
	for key, value := range teamMap { // Order not specified
		fmt.Println(key, value)
	}

	fmt.Println("Configuring the controllers...")

	fmt.Println("setting handlers...")
	//http.Handle("/health-check", controller.DisableCors(&appController))

	fmt.Println("launching server...")
	server := http.Server{
		Addr: ":80",
	}
	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("server launch failed: %v\n", err)
	}


}

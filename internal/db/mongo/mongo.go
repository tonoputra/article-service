package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type dbMongo struct {
	Conn *mongo.Database
	Err  error
}

func (_d *dbMongo) Init() {
	_d.Conn, _d.Err = _d.Connect()
}

// Connect Execute Connect to Mongo server
func (_d *dbMongo) Connect() (*mongo.Database, error) {
	connStr, dbname := _d.ConnStr()

	// Set client options
	clientOptions := options.Client().ApplyURI(connStr)

	// Connect to MongoDB
	dbTimeout := 10 * time.Second // Default

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	connTimeout := 2 * time.Second // Default

	ctx, cancel = context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed when connecting to: %s, error: %v", connStr, err)
	}

	log.Println("Connected to MongoDB: ", os.Getenv("MONGO_HOST"))

	return client.Database(dbname), nil
}

// ConnStr Generate connection string for mongo from Env vars
func (*dbMongo) ConnStr() (string, string) {
	var (
		host     = os.Getenv("MONGO_HOST")
		user     = os.Getenv("MONGO_USERNAME")
		dbname   = os.Getenv("MONGO_DBNAME")
		password = url.QueryEscape(os.Getenv("MONGO_PASSWORD"))
		authdb   = os.Getenv("MONGO_AUTH")
		port     = os.Getenv("MONGO_PORT")
	)

	var connStr string
	if strings.ToLower(host) == "localhost" {
		connStr = fmt.Sprintf("mongodb://%s:%s/%s", host, port, dbname)
	} else {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", user, password, host, port, dbname, authdb)
	}

	return connStr, dbname
}

var Mongo = &dbMongo{}

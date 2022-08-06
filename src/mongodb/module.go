package mongodb

import (
	"context"
	"fmt"
	"github.com/alecthomas/kingpin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoModule struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

func (m *MongoModule) Configure() {
	kingpin.Flag("mongo-host", "Mongo db host").
		Default("localhost").
		Envar("MONGO_HOST").
		StringVar(&m.Host)
	kingpin.Flag("mongo-port", "Mongo db port").
		Default("27017").
		Envar("MONGO_PORT").
		StringVar(&m.Port)
	kingpin.Flag("mongo-user", "Mongo db user").
		Default("root").
		Envar("MONGO_USER").
		StringVar(&m.User)
	kingpin.Flag("mongo-password", "Mongo db password").
		Default("root").
		Envar("MONGO_PASSWORD").
		StringVar(&m.Password)
	kingpin.Flag("mongo-database", "Mongo database name").
		Default("root").
		Envar("MONGO_DATABASE").
		StringVar(&m.Database)
	kingpin.Parse()
}

type MongoDB struct {
	Database    string
	MongoClient *mongo.Client
}

func (m *MongoModule) ProvideMongoDB() (*MongoDB, error) {
	m.Configure()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%v/",
		m.User,
		m.Password,
		m.Host,
		m.Port)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &MongoDB{
		Database:    m.Database,
		MongoClient: client,
	}, nil
}

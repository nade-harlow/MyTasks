package mongoDb

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
	"log"
)

type MongoRepository struct {
	Client *mgo.Database
	DbName string
}

const mongoServerURL = "mongodb://localhost:27017"

//Create a Mongo Client
func NewMongClient() (*mgo.Database, error) {
	dbname := "MyTask"
	DBSession, err := mgo.Dial(mongoServerURL)
	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to Mongo database"))
	}
	log.Println("connected to database successfully")

	Db := DBSession.DB(dbname)
	return Db, nil
}

//NewMongoRepository ...
func (r *MongoRepository) Init() {
	mongoClient, err := NewMongClient()
	r.Client = mongoClient
	r.DbName = "MyTask"
	if err != nil {
		log.Println(err.Error())
	}

}

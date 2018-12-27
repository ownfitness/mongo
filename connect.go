package mongo

import (
	"log"

	"github.com/globalsign/mgo"
)

// Mongo struct to init the connection and set the collection
// database and server
type Mongo struct {
	Server     string
	Database   string
	Collection string
}

// Set db variable for usage in the package
var db *mgo.Database

// Connect to mongo server
func (u *Mongo) Connect(collection string) {
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(u.Database)
}

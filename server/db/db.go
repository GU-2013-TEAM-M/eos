package db

import (
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

// a database instance, to be used outside from this package
var Db *mgo.Database

// connect to the host and select a database
func Connect() {
    mongoSession, err := mgo.Dial("localhost")
    if err != nil {
	panic(err)
    }

    Db = mongoSession.DB("exampledb")
}

// returns a collection for a given name
// caution has to be used, because MongoDB will create
// a new collection automatically, if you specify a new name
func C(name string) *mgo.Collection {
    return Db.C(name)
}

// FIXME: delete alltogether
// it is only here to show that things are working
func Test() string {
    c := C("messages")

    result := Message{}
    err := c.Find(bson.M{}).One(&result)
    if err != nil {
        panic(err)
    }

    return "message: " + result.Msg + " " + result.Id.String()
}

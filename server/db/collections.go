package db

import "labix.org/v2/mgo/bson"

type Message struct {
    Id bson.ObjectId `bson:"_id"`
    Msg string
}

type User struct {
    Email string
    Password string
}

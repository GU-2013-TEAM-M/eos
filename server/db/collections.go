package db

import "labix.org/v2/mgo/bson"

type Message struct {
    Id bson.ObjectId `bson:"_id"`
    Msg string
}

type Organisation struct {
    Id bson.ObjectId `bson:"_id"`
    Name string
}

type User struct {
    Id bson.ObjectId `bson:"_id"`
    OrgId bson.ObjectId
    Email string
    Password string
}

type Daemon struct {
    Id bson.ObjectId `bson:"_id"`
    OrgId bson.ObjectId
    Name string
    Status int
}

type Data struct {
    Id bson.ObjectId `bson:"_id"`
    Type string
    Period string
    Values []float64
}

type Report struct {
    Id bson.ObjectId `bson:"_id"`
    Type string
    Period string
    Values []float64
}

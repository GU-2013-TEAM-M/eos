package db

import "labix.org/v2/mgo/bson"

type Entry interface {
   GenId()
   GetId() bson.ObjectId
}

type Message struct {
    Id bson.ObjectId `bson:"_id"`
    Msg string
}
func (e *Message) GenId() { e.Id = bson.NewObjectId() }
func (e *Message) GetId() bson.ObjectId { return e.Id }

type Organisation struct {
    Id bson.ObjectId `bson:"_id"`
    Name string
}
func (e *Organisation) GenId() { e.Id = bson.NewObjectId() }
func (e *Organisation) GetId() bson.ObjectId { return e.Id }

type Session struct {
    Id bson.ObjectId `bson:"_id"`
    UId bson.ObjectId `bson:"uid"`
}
func (e *Session) GenId() { e.Id = bson.NewObjectId() }
func (e *Session) GetId() bson.ObjectId { return e.Id }

type User struct {
    Id bson.ObjectId `bson:"_id,omitempty"`
    OrgId bson.ObjectId `bson:"org_id"`
    Email string `bson:"email"`
    Password string `bson:"password"`
}
func (e *User) GenId() { e.Id = bson.NewObjectId() }
func (e *User) GetId() bson.ObjectId { return e.Id }

type Daemon struct {
    Id bson.ObjectId `bson:"_id"`
    OrgId bson.ObjectId `bson:"org_id"`
    Password string `bson:"password"`
    Name string `bson:"name"`
    Status string `bson:"status"`
}
func (e *Daemon) GenId() { e.Id = bson.NewObjectId() }
func (e *Daemon) GetId() bson.ObjectId { return e.Id }

type Data struct {
    Id bson.ObjectId `bson:"_id"`
    Type string `bson:"type"`
    Period string `bson:"period"`
    Values []float64 `bson:"values"`
}
func (e *Data) GenId() { e.Id = bson.NewObjectId() }
func (e *Data) GetId() bson.ObjectId { return e.Id }

type Report struct {
    Id bson.ObjectId `bson:"_id"`
    Type string `bson:"type"`
    Period string `bson:"period"`
    Values []float64 `bson:"values"`
}
func (e *Report) GenId() { e.Id = bson.NewObjectId() }
func (e *Report) GetId() bson.ObjectId { return e.Id }

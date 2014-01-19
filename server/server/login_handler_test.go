package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

func Test_userLoginHandler(t *testing.T) {
    tmpU := &db.User{Id: bson.ObjectIdHex("52a4ed348350a921bd000001"),
    OrgId: bson.ObjectIdHex("52a4ed348350a921bd000002"), Email: "a@a.ua", Password: "b"}
    db.AddTemp("users", tmpU)

    // no password
    lcmd := &CmdMessage{Data: make(map[string]interface{}), Conn: &Connection{owner: &User{}}}
    lcmd.Data["email"] = "a@a.ua"

    err := LoginHandler(lcmd)

    test.Assert(err != nil, "it sends an error if there is no password", t)

    // wrong password
    lcmd.Data["password"] = "wrongpass"

    err = LoginHandler(lcmd)

    test.Assert(err != nil, "it sends an error if password is wrong", t)

    // all correct
    lcmd.Data["password"] = "b"

    err = LoginHandler(lcmd)

    cmd := GetLastCmd()
    test.Assert(len(cmd.Data["session_id"].(string)) == 24, "it returns a session_id for this user", t)
    test.Assert(err == nil, "it does not throw an error", t)
    test.Assert(cmd.Data["id"].(string) == "52a4ed348350a921bd000001", "it returns a user id", t)
    session := &db.Session{}
    err = db.C("sessions").Find(bson.M{"_id": bson.ObjectIdHex(cmd.Data["session_id"].(string))}).One(session)
    test.Assert(err == nil, "it creates a session in the database", t)

    db.DelTemps("users")
    db.C("sessions").RemoveAll(bson.M{"uid": bson.ObjectIdHex("52a4ed348350a921bd000001")})

    // no user at all
    err = LoginHandler(lcmd)

    test.Assert(err != nil, "it sends an error if there is no such user", t)
}

func Test_daemonLoginHandler(t *testing.T) {
    // when there is a daemon -----------------------------------------
    tmpD := &db.Daemon{Id: bson.ObjectIdHex("52a4ed348350a921bd000001"),
    OrgId: bson.ObjectIdHex("52a4ed348350a921bd000002"), Name: "daemon", Password: "b"}
    db.AddTemp("daemons", tmpD)

    // a field missing
    lcmd := &CmdMessage{Data: make(map[string]interface{}), Conn: &Connection{owner: &Daemon{}}}
    lcmd.Data["org_id"] = "52a4ed348350a921bd000002"
    lcmd.Data["name"] = "daemon"

    err := LoginHandler(lcmd)

    test.Assert(err != nil, "it sends an error if there is a field missing", t)

    // wrong password
    lcmd.Data["password"] = "wrongpass"

    err = LoginHandler(lcmd)

    test.Assert(err != nil, "it sends an error if password is wrong", t)

    // all correct
    lcmd.Data["password"] = "b"

    err = LoginHandler(lcmd)

    cmd := GetLastCmd()
    test.Assert(cmd.Data["id"].(string) == "52a4ed348350a921bd000001", "it returns a daemon id", t)
    test.Assert(err == nil, "it does not throw an error", t)

    db.DelTemps("daemons")

    // when there is no daemon ----------------------------------------
    err = LoginHandler(lcmd)

    test.Assert(err == nil, "it does not throw an error", t)
    cmd = GetLastCmd()
    err = db.C("daemons").FindId(bson.ObjectIdHex(cmd.Data["id"].(string))).One(tmpD)
    test.Assert(err == nil, "it creates a new daemon, if one does not exist", t)
    test.Assert(tmpD.Status == "NOT_KNOWN", "the daemon is set as unknown", t)

    db.C("daemons").RemoveId(bson.ObjectIdHex(cmd.Data["id"].(string)))
}


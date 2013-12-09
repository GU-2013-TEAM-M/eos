package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

func Test_LoginCheckHandler(t *testing.T) {
    tmp := &db.Session{UId: bson.ObjectIdHex("52a4ed348350a921bd000001")}
    db.AddTemp("sessions", tmp)
    tmpU := &db.User{Id: tmp.UId, OrgId: bson.ObjectIdHex("52a4ed348350a921bd000002"), Email: "a", Password: "b"}
    db.AddTemp("users", tmpU)

    lcmd := &CmdMessage{Data: make(map[string]interface{}), Conn: &Connection{owner: &User{}}}
    lcmd.Data["session_id"] = tmp.Id.Hex()

    LoginCheckHandler(lcmd)

    cmd := GetLastCmd()
    test.Assert(cmd.Data["status"].(string) == "OK", "it recognises the previous session", t)

    db.DelTemps("sessions")
    db.DelTemps("users")

    LoginCheckHandler(lcmd)

    cmd = GetLastCmd()
    test.Assert(cmd.Data["status"].(string) == "UNAUTHORIZED", "it does not authorise user when there is no previous session", t)
}

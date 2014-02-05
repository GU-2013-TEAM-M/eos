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

    msg := &Message{
        msg: `{"type":"loginCheck","data":{
            "session_id": "` + tmp.Id.Hex() + `"
        }}`,
        c: &Connection{ owner: &User{} },
    }

    err, _ := HandleMsg(msg)

    cmd := GetLastCmd()
    test.Assert(cmd.Data["status"].(string) == "OK", "it recognises the previous session", t)

    db.DelTemps("sessions")
    db.DelTemps("users")

    HandleMsg(msg)

    cmd = GetLastCmd()
    test.Assert(cmd.Data["status"].(string) == "UNAUTHORIZED", "it does not authorise user when there is no previous session", t)

    msg.msg = `{"type":"loginCheck","data":{"session_id": "invalid"}}`
    err, _ = HandleMsg(msg)
    test.Assert(err != nil, "It returns an error if session id is invalid objectid", t)
    msg.msg = `{"type":"loginCheck","data":{"session_id": 5}}`
    err, _ = HandleMsg(msg)
    test.Assert(err != nil, "It returns an error if session id is invalid string", t)
}

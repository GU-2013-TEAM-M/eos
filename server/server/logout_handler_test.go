package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

func Test_LogoutHandler(t *testing.T) {
    // a user is removed from the organisation
    // and his session is deleted
    tmpS := &db.Session{UId: bson.ObjectIdHex("52a4ed348350a921bd000001")}
    db.AddTemp("sessions", tmpS)

    user := &User{
        Id: "52a4ed348350a921bd000001",
        SessionId: tmpS.GetId().Hex(),
        OrgId: "Anonymous",
    }
    msg := &Message{
        msg: `{"type":"logout","data":{}}`,
        c: &Connection{ owner: user },
    }
    user.Authorise()

    test.Assert(user.IsAuthorised(), "user is authorised before logout", t)

    err, _ := HandleMsg(msg)

    test.Assert(err == nil, "it logs out successfully", t)
    test.Assert(!user.IsAuthorised(), "user is not in the organisation any more", t)
    err = db.C("sessions").Find(bson.M{
        "uid": bson.ObjectIdHex("52a4ed348350a921bd000001"),
    }).One(tmpS)
    test.Assert(err != nil, "there are no sessions for this user", t)

    // it does not break if session was already gone
    err, _ = HandleMsg(msg)
    test.Assert(err == nil, "it does not blow up without session", t)

    // a daemon is removed from the organisation
    daemon := &Daemon{
        Id: "52a4ed348350a921bd000001",
        OrgId: "Anonymous",
    }
    msg = &Message{
        msg: `{"type":"logout","data":{}}`,
        c: &Connection{ owner: daemon },
    }
    daemon.Authorise()

    test.Assert(daemon.IsAuthorised(), "daemon is authorised before logout", t)

    err, _ = HandleMsg(msg)

    test.Assert(err == nil, "it logs out successfully", t)
    test.Assert(!daemon.IsAuthorised(), "daemon is not in the organisation any more", t)
}


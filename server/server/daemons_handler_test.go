package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
)

func Test_DaemonsHandler(t *testing.T) {
    user := &User{
        Id: "52a4ed348350a921bd000001",
        OrgId: NO_ORG,
    }
    msg := &Message{
        msg: `{"type":"daemons","data":{}}`,
        c: &Connection{ owner: user },
    }

    // user has to be authorised, to get daemons data
    err, _ := HandleMsg(msg)
    test.Assert(err != nil, "user has to be authorised", t)

    // when there are no daemons, it returns an empty list
    user.OrgId = "Anonymous"
    user.Authorise()

    err, _ = HandleMsg(msg)
    test.Assert(err == nil, "it does not throw an error for authorised user", t)

    cmd := GetLastCmd()
    test.Assert(len(cmd.Data["list"].([]map[string]interface{})) == 0, "it does not return inexistent daemons", t)

    // otherwise, it returns a list of daemons
    d1 := &Daemon{Id: "a", OrgId: "Anonymous", Entry: &db.Daemon{}}
    d2 := &Daemon{Id: "b", OrgId: "Anonymous", Entry: &db.Daemon{}}
    d3 := &Daemon{Id: "c", OrgId: "Another_org", Entry: &db.Daemon{}}
    // FIXME: these circular references look bad
    // we should probably refactor them away later on
    d1.c = &Connection{owner: d1}
    d2.c = &Connection{owner: d2}
    d3.c = &Connection{owner: d3}
    d1.Authorise(); d2.Authorise(); d3.Authorise()

    err, _ = HandleMsg(msg)

    cmd = GetLastCmd()
    test.Assert(len(cmd.Data["list"].([]map[string]interface{})) == 2, "it returns only daemons, that are in the same org, as user", t)

    // daemons cannot request it
    msg.c = &Connection{owner: &Daemon{}}

    err, _ = HandleMsg(msg)
    test.Assert(err != nil, "daemons are disallowed", t)

    // cleaning up
    user.Deauthorise()
    d1.Deauthorise()
    d2.Deauthorise()
    d3.Deauthorise()
}


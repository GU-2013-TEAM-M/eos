package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
)

func Test_ControlHandler(t *testing.T) {
    user := &User{
        Id: "52a4ed348350a921bd000001",
        OrgId: NO_ORG,
    }
    msg := &Message{
        msg: `{
            "type": "control",
            "data": {
                "daemon_id": "a",
                "operation": {
                    "start": ["cpu"]
                }
            }
        }`,
        c: &Connection{ owner: user },
    }
    daemon := &Daemon{
        Id: "a",
        OrgId: "Random_org",
        c: &Connection{},
        Entry: &db.Daemon {
            Parameters: []string { "cpu", "ram" },
        },
    }
    daemon.c.owner = daemon;
    daemon.Authorise()

    // the user has to be authorised
    err, _ := HandleMsg(msg)
    test.Assert(err != nil, "user has to be authorised", t)

    // it fails if the daemon does not exist
    user.OrgId = "Anonymous"
    user.Authorise()

    err, _ = HandleMsg(msg)
    test.Assert(err != nil, "the daemon has to exist in the org", t)

    // it transmutes the correct message if the daemon exists
    daemon.Deauthorise()
    daemon.OrgId = "Anonymous"
    daemon.Authorise()

    err, _ = HandleMsg(msg)
    test.Assert(err == nil, "it does not send the error, when all is ok", t)

    // but it throws an error if we ask for unsupported parameters
    msg.msg = `{
        "type": "control",
        "data": {
            "daemon_id": "a",
            "operation": {
                "start": ["cpu"],
                "stop": ["something"]
            }
        }
    }`

    err, _ = HandleMsg(msg)
    test.Assert(err != nil, "it does not allow to ask for unsupported parameters", t)

    // it is only available to user
    msg.c = &Connection{owner: &Daemon{}}

    err, _ = HandleMsg(msg)
    test.Assert(err != nil, "daemons are disallowed", t)

    // cleaning up
    daemon.Deauthorise()
    user.Deauthorise()
}


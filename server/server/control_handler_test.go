package main

import (
    "testing"
    "eos/server/test"
)

func Test_ControlHandler(t *testing.T) {
    user := &User{
        Id: "52a4ed348350a921bd000001",
        OrgId: NO_ORG,
    }
    lcmd := &CmdMessage{
        Data: make(map[string]interface{}),
        Conn: &Connection{owner: user},
    }
    lcmd.Data["daemon_id"] = "a"
    daemon := &Daemon{
        Id: "a",
        OrgId: "Random_org",
        c: &Connection{},
    }
    daemon.Authorise()

    // the user has to be authorised
    err := ControlHandler(lcmd)
    test.Assert(err != nil, "user has to be authorised", t)

    // it fails if the daemon does not exist
    user.OrgId = "Anonymous"
    user.Authorise()

    err = ControlHandler(lcmd)
    test.Assert(err != nil, "the daemon has to exist in the org", t)

    // it transmutes the message if the daemon exists
    daemon.Deauthorise()
    daemon.OrgId = "Anonymous"
    daemon.Authorise()

    err = ControlHandler(lcmd)
    test.Assert(err == nil, "it does not send the error, when all is ok", t)

    // it is only available to user
    lcmd.Conn = &Connection{owner: &Daemon{}}

    err = ControlHandler(lcmd)
    test.Assert(err != nil, "daemons are disallowed", t)

    // cleaning up
    daemon.Deauthorise()
    user.Deauthorise()
}


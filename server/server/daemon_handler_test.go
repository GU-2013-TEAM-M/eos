package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
)

func Test_DaemonHandler(t *testing.T) {
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
        Entry: &db.Daemon{},
    }
    daemon.c.owner = daemon
    daemon.Authorise()

    // user has to be authorised, to get daemons data
    err := DaemonHandler(lcmd)
    test.Assert(err != nil, "user has to be authorised", t)

    // if the daemon does not exist, then an error is returned
    user.OrgId = "Anonymous"
    user.Authorise()

    err = DaemonHandler(lcmd)
    test.Assert(err != nil, "the daemon has to exist in the org", t)

    // if everything is ok, the daemon information is returned
    daemon.Deauthorise()
    daemon.OrgId = "Anonymous"
    daemon.Authorise()

    err = DaemonHandler(lcmd)
    test.Assert(err == nil, "it does not send the error", t)

    cmd := GetLastCmd()
    test.Assert(cmd.Data["daemon_id"].(string) == "a", "it returns the daemon information", t)

    // daemons cannot request it
    lcmd.Conn = &Connection{owner: &Daemon{}}

    err = DaemonsHandler(lcmd)
    test.Assert(err != nil, "daemons are disallowed", t)

    // cleaning up
    daemon.Deauthorise()
    user.Deauthorise()
}


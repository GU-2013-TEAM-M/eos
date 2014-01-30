package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
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

    //--------------------------------------------------
    // if it is a daemon
    //--------------------------------------------------
    // not sending the data
    lcmd.Conn = &Connection{owner: daemon}

    err = DaemonHandler(lcmd)
    test.Assert(err != nil, "it has to contain information", t)

    // when the data is sent, it stores it in the database
    lcmd.Data["daemon_platform"] = "Linux"
    lcmd.Data["daemon_all_parameters"] = []string{"cpu", "ram", "network"}
    lcmd.Data["daemon_monitored_parameters"] = []string{"cpu", "ram"}

    tmpD := &db.Daemon{OrgId: bson.ObjectIdHex("52a4ed348350a921bd000002"), Name: "a", Password: "b"}
    db.AddTemp("daemons", tmpD)
    daemon.Entry = tmpD
    daemon.Id = tmpD.Id.Hex()

    err = DaemonHandler(lcmd)
    test.Assert(err == nil, "no errors are raised", t)
    dbd := &db.Daemon{}
    db.C("daemons").FindId(tmpD.Id).One(dbd)
    test.Assert(dbd.Platform == lcmd.Data["daemon_platform"], "it stores the platform in the database", t)
    test.Assert(len(dbd.Parameters) == len(lcmd.Data["daemon_all_parameters"].([]string)), "it stores all the parameters", t)
    test.Assert(len(dbd.Monitored) == len(lcmd.Data["daemon_monitored_parameters"].([]string)), "it also stores what it is monitoring", t)

    // it also sends new information to all the users in the org
    cmd = GetLastCmd()
    test.Assert(cmd.Type == "daemon", "it sends a daemon message", t)
    test.Assert(cmd.Data["daemon_id"].(string) == daemon.Id, "with the latest information about this daemon", t)

    // cleaning up
    daemon.Deauthorise()
    user.Deauthorise()
    db.DelTemps("daemons")
}


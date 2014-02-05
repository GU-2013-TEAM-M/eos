package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

func Test_MonitoringHandlerDaemon(t *testing.T) {
    daemon := &Daemon{Id: "a", OrgId: NO_ORG, Entry: &db.Daemon{}}

    lcmd := &CmdMessage{
        Data: make(map[string]interface{}),
        Conn: &Connection{owner: daemon},
    }
    value := make(map[int64]float64)
    value[1104699] = 12.5
    entry := make(map[string]interface{})
    entry["parameter"] = "cpu"
    entry["values"] = value
    lcmd.Data["list"] = []interface{}{entry}

    // without authorisation
    err := MonitoringHandler(lcmd)

    test.Assert(err != nil, "the daemon must be authorised", t)

    // with authorisation
    daemon.OrgId = "Anonymous"
    daemon.Authorise()

    err = MonitoringHandler(lcmd)

    test.Assert(err == nil, "it does not throw errors, when the daemon is authorised", t)
    data := &db.Data{}
    db.C("monitoring_of_a").Find(bson.M{"parameter": "cpu"}).One(data)
    test.Assert(data.Time == 1104699, "it stores the time right", t)
    test.Assert(data.Value == 12.5, "it stores the metric right", t)

    // adding additional metrics to it
    value2 := make(map[int64]float64)
    value2[1104670] = 15
    entry["values"] = value2

    err = MonitoringHandler(lcmd)

    test.Assert(err == nil, "it still doesn't throw errors", t)
    q := db.C("monitoring_of_a").Find(bson.M{"parameter": "cpu"})
    count, err := q.Count()
    test.Assert(count == 2, "it returns additional values", t)
    q.Sort("time").One(data)
    test.Assert(data.Time == 1104670, "it stores the time right", t)
    test.Assert(data.Value == 15, "it stores the metric right", t)

    // cleaning up
    db.C("monitoring_of_a").DropCollection()
    daemon.Deauthorise()
}

func Test_MonitoringHandlerUser(t *testing.T) {
}

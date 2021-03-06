package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

func Test_MonitoringHandlerDaemon(t *testing.T) {
    daemon := &Daemon{Id: "a", OrgId: NO_ORG, Entry: &db.Daemon{}}

    msg := &Message{
        msg: `{
            "type": "monitoring",
            "data": {
                "list": [{
                    "parameter": "cpu",
                    "values": {
                        "1104699": 12.5
                    }
                }]
            }
        }`,
        c: &Connection{ owner: daemon },
    }

    // without authorisation
    err, _ := HandleMsg(msg)

    test.Assert(err != nil, "the daemon must be authorised", t)

    // with authorisation
    daemon.OrgId = "Anonymous"
    daemon.Authorise()

    err, _ = HandleMsg(msg)

    test.Assert(err == nil, "it does not throw errors, when the daemon is authorised", t)
    data := &db.Data{}
    db.C("monitoring_of_a").Find(bson.M{"parameter": "cpu"}).One(data)
    test.Assert(data.Time == 1104699, "it stores the time right", t)
    test.Assert(data.Value == 12.5, "it stores the metric right", t)

    // adding additional metrics to it
    msg.msg = `{
        "type": "monitoring",
        "data": {
            "list": [{
                "parameter": "cpu",
                "values": {
                    "1104670": 15
                }
            }]
        }
    }`

    err, _ = HandleMsg(msg)

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
    // before
    user := &User{ OrgId: "Anonymous" }
    daemon := &Daemon{ Id: "a", OrgId: "Anonymous" }
    user.Authorise()

    data1 := &db.Data{ "", "cpu", 1000, 12.5 }
    data2 := &db.Data{ "", "cpu", 1500, 14.5 }
    data3 := &db.Data{ "", "cpu", 1600, 15.5 }
    data4 := &db.Data{ "", "cpu", 1900, 11.5 }
    data5 := &db.Data{ "", "ram", 1200, 9000 }
    db.AddTemp( "monitoring_of_a", data1)
    db.AddTemp( "monitoring_of_a", data2)
    db.AddTemp( "monitoring_of_a", data3)
    db.AddTemp( "monitoring_of_a", data4)
    db.AddTemp( "monitoring_of_a", data5)

    // let's try it from the string...
    msg := &Message{
        msg: `
        {
            "type": "monitoring",
            "data": {
                "daemon_id": "a",
                "parameter": "cpu",
                "from": 1100,
                "to": 1600
            }
        }
        `,
        c: &Connection{ owner: user },
    }

    // the daemon is not in the org
    err, _ := HandleMsg(msg)

    test.Assert(err != nil, "it doesn't allow to monitor foreign daemons", t)

    // daemon exists in the org
    daemon.Authorise()

    err, _ = HandleMsg(msg)

    cmd := GetLastCmd()

    test.Assert(err == nil, "it does allow to monitor your daemons", t)
    vals := cmd.Data["values"].(map[string]float64)
    test.Assert(len(vals) == 2, "it returns the right number of answers", t)
    test.Assert(vals["1500"] == 14.5, "it has the correct data", t)
    test.Assert(vals["1600"] == 15.5, "it has the correct data", t)

    // cleaning up
    db.C("monitoring_of_a").DropCollection()
    user.Deauthorise(); daemon.Deauthorise()
}

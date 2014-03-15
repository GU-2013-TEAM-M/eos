package main

import (
    "fmt"
    "strconv"
    "errors"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

// a handler that either stores monitoring information
// for that particular daemon, or sends the stats to the user
func MonitoringHandler(cmd *CmdMessage) error {
    fmt.Println("I understood you wanted history")
    if !cmd.Conn.owner.IsAuthorised() {
        return errors.New("Logging in is required before using this handler")
    }
    if !cmd.Conn.owner.IsUser() {
        return storeMonitoringData(cmd)
    }
    fmt.Println("I understood you are the user")

    daemonId, ok := cmd.Data["daemon_id"].(string)
    if !ok {
        return errors.New("wrong daemon_id supplied")
    }

    daemons := cmd.Conn.owner.GetOrg().Daemons
    _, ok = daemons[daemonId]
    if !ok {
        return errors.New("daemon not found")
    }

    param, ok1 := cmd.Data["parameter"].(string)
    from, ok2 := cmd.Data["from"].(float64)
    to, ok3 := cmd.Data["to"].(float64)

    if (!(ok1 && ok2 && ok3)) {
        return errors.New("Some parameters are missing!")
    }

    c := db.C("monitoring_of_" + daemonId);
    var inf []db.Data = nil;

    c.Find(bson.M{
        "time": bson.M{ "$gte": from, "$lte": to },
        "parameter": param,
    }).All(&inf)

    vals := make(map[string]float64)

    for _, d := range inf {
        vals[strconv.FormatInt(d.Time, 10)] = d.Value
    }

    data := make(map[string]interface{})
    data["daemon_id"] = daemonId
    data["parameter"] = param
    data["values"] = vals

    err := DispatchMessage("monitoring", data, cmd.Conn)
    fmt.Println(err)
    fmt.Println("And I did send you all the data")

    return nil
}

func storeMonitoringData(cmd *CmdMessage) error {
    // creating a new data point
    list, ok := cmd.Data["list"].([]interface{})
    if !ok {
        return errors.New("A list of parameters is not specified")
    }

    c := db.C("monitoring_of_" + cmd.Conn.owner.(*Daemon).Id)
    for _, e := range list {
        entry, ok1 := e.(map[string]interface{})
        if !ok1 {
            return errors.New("Monitoring entry is invalid")
        }
        parameter, ok2 := entry["parameter"].(string)
        values, ok3 := entry["values"].(map[string]interface{})
        if !(ok2 && ok3) {
            return errors.New("You need to specify parameter and values")
        }

        for t, v := range values {
            time, _ := strconv.Atoi(t)
            value, ok := v.(float64)
            if !ok {
                return errors.New("Non floating point number value supplied")
            }
            c.Insert(bson.M{
                "parameter": parameter,
                "time": time,
                "value": value,
            })
        }
    }

    return nil
}

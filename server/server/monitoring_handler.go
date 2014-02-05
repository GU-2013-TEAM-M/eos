package main

import (
    "errors"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

// a handler that either stores monitoring information
// for that particular daemon, or sends the stats to the user
func MonitoringHandler(cmd *CmdMessage) error {
    if !cmd.Conn.owner.IsAuthorised() {
        return errors.New("Logging in is required before using this handler")
    }
    if !cmd.Conn.owner.IsUser() {
        return storeMonitoringData(cmd)
    }

    return errors.New("not implemented")
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
        values, ok3 := entry["values"].(map[int64]float64)
        if !(ok2 && ok3) {
            return errors.New("You need to specify parameter and values")
        }

        for time, value := range values {
            c.Insert(bson.M{
                "parameter": parameter,
                "time": time,
                "value": value,
            })
        }
    }

    return nil
}

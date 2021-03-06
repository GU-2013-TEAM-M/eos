package main

import (
    "errors"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

// a handler that provides the detailed information about the daemon
func DaemonHandler(cmd *CmdMessage) error {
    if !cmd.Conn.owner.IsAuthorised() {
        return errors.New("Logging in is required before using this handler")
    }
    if !cmd.Conn.owner.IsUser() {
        return daemonUpdateInfo(cmd)
    }

    daemonId, ok := cmd.Data["daemon_id"].(string)
    if !ok {
        return errors.New("wrong daemon_id supplied")
    }

    daemons := cmd.Conn.owner.GetOrg().Daemons
    daemon, ok := daemons[daemonId]
    if !ok {
        return errors.New("daemon not found: " + daemonId)
    }

    sendDaemonMessage(daemon.owner.(*Daemon), cmd.Conn)
    return nil
}

func daemonUpdateInfo(cmd *CmdMessage) error {
    platform, ok1 := cmd.Data["daemon_platform"].(string)
    par, ok2 := cmd.Data["daemon_all_parameters"].([]interface{})
    mon, ok3 := cmd.Data["daemon_monitored_parameters"].([]interface{})

    if !(ok1 && ok2 && ok3) {
        return errors.New("All parameters must be supplied")
    }

    parameters, err1 := ToStringSlice(par)
    monitored, err2 := ToStringSlice(mon)

    if err1 != nil || err2 != nil {
        return errors.New("Non-string parameters given")
    }

    // update elements for this id
    daemon := cmd.Conn.owner.(*Daemon)
    db.C("daemons").UpdateId(daemon.Entry.Id, bson.M{
        "$set": bson.M{
            "platform": platform,
            "parameters": parameters,
            "monitored": monitored,
        },
    })

    // update them on the daemon itself
    daemon.Entry.Platform = platform
    daemon.Entry.Parameters = parameters
    daemon.Entry.Monitored = monitored

    // send information to all the users in the org
    for _, ucon := range daemon.GetOrg().Users {
        err := sendDaemonMessage(daemon, ucon)
        if err != nil {
            return err
        }
    }

    return nil
}

func sendDaemonMessage(daemon *Daemon, c *Connection) error {
    data := make(map[string]interface{})
    data["daemon_id"] = daemon.Id

    // I cannot stub websockets properly, hence, for test purposes...
    ip := "127.0.0.1:8080"
    if daemon.c.ws != nil {
        ip = daemon.c.ws.Request().RemoteAddr
    }
    data["daemon_address"] = "ws://" + ip

    d := daemon.Entry
    data["daemon_platform"] = d.Platform
    data["daemon_all_parameters"] = d.Parameters
    data["daemon_monitored_parameters"] = d.Monitored

    return DispatchMessage("daemon", data, c)
}

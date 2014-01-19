package main

import (
    "errors"
)

// a handler that provides the detailed information about the daemon
func DaemonHandler(cmd *CmdMessage) error {
    if !cmd.Conn.owner.IsUser() {
        return errors.New("This handler is not available for daemons")
    }
    if !cmd.Conn.owner.IsAuthorised() {
        return errors.New("User has to log in before using this handler")
    }

    //sessId := cmd.Data["session_id"].(string)
    daemonId, ok := cmd.Data["daemon_id"].(string)
    if !ok {
        return errors.New("wrong daemon_id supplied")
    }

    data := make(map[string]interface{})

    daemons := cmd.Conn.owner.GetOrg().Daemons
    daemon, ok := daemons[daemonId]
    if !ok {
        return errors.New("daemon not found")
    }

    data["daemon_id"] = daemonId

    // I cannot stub websockets properly, hence, for test purposes...
    ip := "127.0.0.1:8080"
    if daemon.ws != nil {
        ip = daemon.ws.Request().RemoteAddr
    }
    data["daemon_address"] = "ws://" + ip

    d := daemon.owner.(*Daemon).Entry
    data["daemon_platform"] = d.Platform
    data["daemon_all_parameters"] = d.Parameters
    data["daemon_monitored_parameters"] = d.Monitored

    DispatchMessage("daemon", data, cmd.Conn)
    return nil
}

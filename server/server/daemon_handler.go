package main

import (
    "errors"
    //"eos/server/db"
)

// a handler that provides the detailed information about the daemon
func DaemonHandler(cmd *CmdMessage) error {
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
    data["daemon_address"] = "ws://" + daemon.ws.Request().RemoteAddr
    data["daemon_platform"] = []string{"Linux"}
    data["daemon_all_parameters"] = []string{"CPU", "Memory"}
    data["daemon_monitored_parameters"] = []string{"CPU"}

    DispatchMessage("daemon", data, cmd.Conn)
    return errors.New("Daemon: Not implemented")
}

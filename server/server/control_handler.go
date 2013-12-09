package main

import (
    "errors"
    //"eos/server/db"
)

// a handler that checks that the current session id is still active
func ControlHandler(cmd *CmdMessage) error {
    //sessId := cmd.Data["session_id"].(string)
    daemonId := cmd.Data["daemon_id"].(string)

    data := make(map[string]interface{})
    data["daemon_id"] = daemonId

    daemon, ok := cmd.Conn.owner.(*User).GetOrg().Daemons[daemonId]
    if !ok {
        data["status"] = "NOT_OK"
        DispatchMessage("control", data, cmd.Conn)
        return errors.New("Trying to control inexistent daemon")
    }

    data["status"] = "OK"
    DispatchMessage("control", data, cmd.Conn)
    DispatchMessage("control", cmd.Data, daemon)
    return nil
}

package main

import (
    "errors"
    //"eos/server/db"
)

// a handler that returns all the daemons we have
func DaemonsHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    data["list"] = [](map[string]interface{}){}
    daemons := cmd.Conn.owner.GetOrg().Daemons

    for id := range daemons {
        data["list"] = append(data["list"].([]map[string]interface{}), getDaemonFormat(daemons[id]))
    }

    DispatchMessage("daemons", data, cmd.Conn)
    return errors.New("Daemons:Daemons: Not implemented")
}

// a helper function, generating the daemon info from the connection
func getDaemonFormat(c *Connection) map[string]interface{} {
    daemon := make(map[string]interface{})
    daemon["daemon_id"] = c.owner.(*Daemon).Id
    daemon["daemon_name"] = "Gabriel's laptop"
    daemon["daemon_state"] = "RUNNING"
    return daemon
}

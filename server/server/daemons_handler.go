package main

import (
    "errors"
)

// a handler that returns all the daemons we have
func DaemonsHandler(cmd *CmdMessage) error {
    if !cmd.Conn.owner.IsUser() {
        return errors.New("This handler is not available for daemons")
    }
    if !cmd.Conn.owner.IsAuthorised() {
        return errors.New("User has to log in before using this handler")
    }

    data := make(map[string]interface{})
    data["list"] = [](map[string]interface{}){}
    daemons := cmd.Conn.owner.GetOrg().Daemons

    for id := range daemons {
        data["list"] = append(data["list"].([]map[string]interface{}), getDaemonFormat(daemons[id]))
    }

    DispatchMessage("daemons", data, cmd.Conn)
    return nil
}

// a helper function, generating the daemon info from the connection
func getDaemonFormat(c *Connection) map[string]interface{} {
    d := c.owner.(*Daemon)
    daemon := make(map[string]interface{})
    daemon["daemon_id"] = d.Id
    daemon["daemon_name"] = d.Entry.Name
    //daemon["daemon_state"] = d.Entry.Status
    daemon["daemon_state"] = "Running"
    return daemon
}

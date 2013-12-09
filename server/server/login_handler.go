package main

import (
    "errors"
    //"eos/server/db"
    "time"
)

// a handler that logs in the user
func LoginHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    data["session_id"] = "52a4ed348350a921bd000001"
    if cmd.Conn.owner.IsUser() {
        cmd.Conn.owner.(*User).Id = string(time.Now().Nanosecond())
        cmd.Conn.owner.(*User).OrgId = "Anonymous"
    } else {
        cmd.Conn.owner.(*Daemon).Id = string(time.Now().Nanosecond())
        cmd.Conn.owner.(*Daemon).OrgId = "Anonymous"
    }
    DispatchMessage("login", data, cmd.Conn)
    return errors.New("Login: Not implemented")
}

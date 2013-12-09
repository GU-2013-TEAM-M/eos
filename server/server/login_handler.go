package main

import (
    "errors"
    //"eos/server/db"
    "time"
    "strconv"
)

// a handler that logs in the user
func LoginHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    data["session_id"] = "52a4ed348350a921bd000001"
    if cmd.Conn.owner.IsUser() {
        cmd.Conn.owner.(*User).Id = strconv.Itoa(time.Now().Nanosecond())
        cmd.Conn.owner.(*User).OrgId = "Anonymous"
        data["id"] = cmd.Conn.owner.(*User).Id
    } else {
        cmd.Conn.owner.(*Daemon).Id = strconv.Itoa(time.Now().Nanosecond())
        cmd.Conn.owner.(*Daemon).Id = "12345"
        cmd.Conn.owner.(*Daemon).OrgId = "Anonymous"
        data["id"] = cmd.Conn.owner.(*Daemon).Id
    }
    cmd.Conn.owner.Authorise()
    DispatchMessage("login", data, cmd.Conn)
    return errors.New("Login: Not implemented")
}

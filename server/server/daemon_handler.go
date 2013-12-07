package main

import (
    "errors"
    //"eos/server/db"
)

// a handler that checks that the current session id is still active
func DaemonHandler(cmd *CmdMessage) error {
    return errors.New("Not implemented")
}

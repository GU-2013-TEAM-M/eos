package main

import (
    "errors"
)

// a handler that checks that the current session id is still active
func ControlHandler(cmd *CmdMessage) error {
    if !cmd.Conn.owner.IsUser() {
        return errors.New("This handler is not available for daemons")
    }
    if !cmd.Conn.owner.IsAuthorised() {
        return errors.New("User has to log in before using this handler")
    }

    daemonId := cmd.Data["daemon_id"].(string)

    data := make(map[string]interface{})
    data["daemon_id"] = daemonId

    daemon, ok := cmd.Conn.owner.(*User).GetOrg().Daemons[daemonId]
    if !ok {
        return errors.New("Trying to control inexistent daemon")
    }

    // doing a pre-check, if the daemon supports the said parameters
    operation, ok := cmd.Data["operation"].(map[string]interface{})
    if !ok {
        return errors.New("Format error")
    }
    start, startGiven := operation["start"].([]interface{})
    stop, stopGiven := operation["stop"].([]interface{})
    parameters := daemon.owner.(*Daemon).Entry.Parameters

    if startGiven && !isSubset(start, parameters) ||
        stopGiven && !isSubset(stop, parameters) {
        return errors.New("Asking to start/stop unsupported parameter")
    }

    data["status"] = "OK"
    DispatchMessage("control", cmd.Data, daemon)
    return DispatchMessage("control", data, cmd.Conn)
}

func isSubset(smaller []interface{}, larger []string) bool {
    for _, el := range smaller {
        str, ok := el.(string)
        if !ok || !inArray(str, larger) {
            return false;
        }
    }
    return true;
}

func inArray(a string, arr []string) bool {
    for _, el := range arr {
        if a == el {
            return true;
        }
    }
    return false;
}

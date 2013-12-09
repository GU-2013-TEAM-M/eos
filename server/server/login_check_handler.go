package main

import (
    "errors"
    "labix.org/v2/mgo/bson"
)

// a handler that checks that the current session id is still active
func LoginCheckHandler(cmd *CmdMessage) error {
    sessId, ok := cmd.Data["session_id"].(string)
    if !ok || !bson.IsObjectIdHex(sessId) {
        return errors.New("Invalid session id supplied")
    }
    data := make(map[string]interface{})

    if cmd.Conn.owner.IsUser() {
        id, err := AuthFromSession(sessId)

        if err == nil {
            // everything is ok, logging in user
            err = cmd.Conn.owner.Authenticate(id)
            if err != nil {
                data["status"] = "UNAUTHORIZED"
            } else {
                // sending back the success message
                data["status"] = "OK"
            }
        } else {
            // sending back the failure message
            data["status"] = "UNAUTHORIZED"
        }

        DispatchMessage("loginCheck", data, cmd.Conn)
        return err
    }

    return errors.New("This handler is only available for users")
}

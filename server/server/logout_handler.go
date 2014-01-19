package main

import (
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

// a handler that deauthorises the user or the daemon
func LogoutHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    cmd.Conn.owner.Deauthorise()

    // if it is user, delete the session
    if cmd.Conn.owner.IsUser() {
        db.C("sessions").RemoveAll(bson.M{
            "uid": bson.ObjectIdHex(cmd.Conn.owner.(*User).Id),
        })
    }

    data["status"] = "OK"
    DispatchMessage("logout", data, cmd.Conn)
    return nil
}
